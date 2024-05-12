pipeline {
    agent any
    stages {
        stage('decrypt prod.env') {
            steps {
                sh "gpg --decrypt ./env/prod.env.asc > prod.env"
                sh "mv ./env/prod.env ./.env"
                sh "cat ./.env"
            }
        }
        stage('build') {
            steps {
                sh "docker build -t docker.alekseikromski.com/atlanta-server:latest ."
            }
        }
        stage('push') {
            steps {
                withCredentials([usernamePassword(credentialsId: 'dockerHub', passwordVariable: 'dockerHubPassword', usernameVariable: 'dockerHubUser')]) {
                    sh "docker login https://docker.alekseikromski.com -u ${env.dockerHubUser} -p ${env.dockerHubPassword}"
                    sh "docker push docker.alekseikromski.com/atlanta-server:latest"
                }
            }
        }
        stage('Apply Kubernetes files') {
            steps {
                withEnv(["KUBECONFIG=/.kube/config"]) {
                    sh "kubectl apply -f ./k8s/deploy.yml --force"
                    sh "kubectl apply -f ./k8s/service.yml --force"
                }
            }
        }
    }
}

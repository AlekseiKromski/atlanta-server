import Wrapper from "../../components/common/wrapper/wrapper";
import SettingsStyle from "./settings.module.css"
import {Button, Input} from "@nextui-org/react";
import {useDispatch, useSelector} from "react-redux";
import {useEffect, useState} from "react";
import Alert from "../../components/ui/alert/alert";
import {setUser as reduxSetUser} from "../../store/application/application";

export default function Settings() {

    const application = useSelector((state) => state.application);
    const dispatch = useDispatch()
    const [alert, setAlert] = useState(null);
    const [user, setUser] = useState(        {
        "username": "",
        "first_name": "",
        "second_name": "",
        "password": "",
        "email": "",
        "role": "",
        "id": "",
    })

    const onUpsert = () => {
        if (
            user.username.length === 0 ||
            user.email.length === 0 ||
            user.first_name.length === 0 ||
            user.second_name.length === 0 ||
            user.role.length === 0
        ) {
            return
        }
        application.axios.post("/api/users/current-user-upsert", {
            "username": user.username,
            "first_name": user.first_name,
            "second_name": user.second_name,
            "password": user.password,
            "email": user.email,
            "role": user.role,
            "id": user.id.length === 0 ? null : user.id,
        })
            .then(e => {
                dispatch(reduxSetUser({
                    ...application.user,
                    "username": user.username,
                    "first_name": user.first_name,
                    "second_name": user.second_name,
                    "email": user.email,
                    "password": "",
                }))
            })
            .catch(e => {
                setAlert({
                    type: "error",
                    description: `Cannot upsert user: ${e}`,
                    link: "https://google.com"
                })
                setTimeout(() => setAlert(null), 5000)
            })
    }

    useEffect(() => {
        if (!application.user) {
            return
        }
        setUser({...application.user, password: ""})
    }, [application.user]);

    return (
        <div className={SettingsStyle.SettingsBody + " w-full flex flex-col"}>
            <Wrapper width="100%" title="Search" modal={{
                title: "Settings",
                body: (
                    <p>
                        TODO: ...
                    </p>
                )
            }}>
                <div className={SettingsStyle.SettingsPanel + " flex flex-col"}>
                    <Input
                        type="username"
                        label="Username"
                        className="w-full"
                        onChange={(e) => {
                            setUser({
                                ...user,
                                username: e.target.value,
                            })
                        }}
                        value={user.username}
                    />

                    <Input
                        type="first_name"
                        label="First name"
                        className="w-full"
                        onChange={(e) => {
                            setUser({
                                ...user,
                                first_name: e.target.value,
                            })
                        }}
                        value={user.first_name}
                    />

                    <Input
                        type="second_name"
                        label="Second name"
                        className="w-full"
                        onChange={(e) => {
                            setUser({
                                ...user,
                                second_name: e.target.value,
                            })
                        }}
                        value={user.second_name}
                    />

                    <Input
                        type="email"
                        label="Email"
                        className="w-full"
                        onChange={(e) => {
                            setUser({
                                ...user,
                                email: e.target.value,
                            })
                        }}
                        value={user.email}
                    />

                    <Input
                        type="password"
                        label="Password"
                        className="w-full"
                        onChange={(e) => {
                            setUser({
                                ...user,
                                password: e.target.value,
                            })
                        }}
                        value={user.password}
                    />

                    <div>
                        <Button onClick={() => {onUpsert()}} variant="flat" color="warning">Update</Button>
                    </div>
                </div>
            </Wrapper>
            {
                alert &&
                <Alert type={alert.type} description={alert.description} link={alert.link}/>
            }
        </div>
    )
}
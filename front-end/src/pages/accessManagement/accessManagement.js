import Wrapper from "../../components/common/wrapper/wrapper";
import AccessManagementStyle from "./accessManagement.module.css"
import Alert from "../../components/ui/alert/alert";
import {useState} from "react";
import {Card, CardBody, Tab, Tabs} from "@nextui-org/react";
import Users from "../../components/accessManagement/users/users"
import Roles from "../../components/accessManagement/roles/roles";
import Endpoints from "../../components/accessManagement/endpoints/endpoints";

export default function AccessManagement() {
    const [alert, setAlert] = useState(null);

    return (
        <div className={AccessManagementStyle.DevicesBody + " w-full"}>
            <Wrapper
                title="Access management"
                modal={
                    {
                        title: "Access management",
                        body: (
                            <p>
                                TODO: ...
                            </p>
                        )
                    }
                }
            >
                <Tabs aria-label="Options">
                    <Tab key="Users" title="Users">
                        <Card>
                            <CardBody>
                                <Users/>
                            </CardBody>
                        </Card>
                    </Tab>
                    <Tab key="roles" title="Roles">
                        <Card>
                            <CardBody>
                                <Roles/>
                            </CardBody>
                        </Card>
                    </Tab>
                    <Tab key="endpoints" title="Endpoints">
                        <Card>
                            <CardBody>
                                <Endpoints/>
                            </CardBody>
                        </Card>
                    </Tab>
                </Tabs>
            </Wrapper>
            {
                alert &&
                <Alert type={alert.type} description={alert.description} link={alert.link}/>
            }
        </div>
    )
}
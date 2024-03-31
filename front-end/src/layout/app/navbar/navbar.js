import NavbarStyles from "./navbar.module.css"
import {Link} from "react-router-dom";
import {Button, Dropdown, DropdownItem, DropdownMenu, DropdownTrigger} from "@nextui-org/react";
import {useSelector} from "react-redux";

export default function Navbar() {

    const application = useSelector((state) => state.application);

    return (
        <div className={NavbarStyles.Navbar + " flex justify-between"}>
            <Link to={"/"}>Atlanta</Link>
            <div class="gap-1.5 flex justify-center items-center">
                <span>Hi, Aleksei Kromski</span>
                <Dropdown>
                    <DropdownTrigger>
                        <Button
                            size="sm"
                            color="primary" variant="ghost"
                        >
                            Menu
                        </Button>
                    </DropdownTrigger>
                    <DropdownMenu aria-label="Static Actions">
                        <DropdownItem key="new">Settings</DropdownItem>
                        <DropdownItem key="delete" className="text-danger" color="danger" onClick={
                            () => {
                                application.axios.get("/api/auth/logout")
                                    .then(res => {
                                        window.location = "/"
                                    })
                                    .catch(e => console.log(e))
                            }
                        }>
                            Logout
                        </DropdownItem>
                    </DropdownMenu>
                </Dropdown>
            </div>
        </div>
    )
}
import NavbarStyles from "./navbar.module.css"
import {Link} from "react-router-dom";
import {Button, Dropdown, DropdownItem, DropdownMenu, DropdownTrigger} from "@nextui-org/react";

export default function Navbar() {
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
                        <DropdownItem key="delete" className="text-danger" color="danger">
                            Logout
                        </DropdownItem>
                    </DropdownMenu>
                </Dropdown>
            </div>
        </div>
    )
}
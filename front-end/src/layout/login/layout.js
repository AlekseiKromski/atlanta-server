import {Outlet} from "react-router-dom"
import LayoutStyle from "./layout.module.css"

function Layout() {
    return (
        <div className={LayoutStyle.LoginRoot + " flex flex-col h-full justify-center items-center dark text-foreground bg-background"}>
            <div className={LayoutStyle.Login + " flex flex-col"}>
                <h1>Welcome to Atlanta</h1>
                <Outlet/>
            </div>
        </div>
    );
}

export default Layout;

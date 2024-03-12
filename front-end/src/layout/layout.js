import {Outlet} from "react-router-dom"
import Sidebar from "./sidebar/sidebar";
import Navbar from "./navbar/navbar";

function Layout() {
    return (
        <div class="flex flex-col h-full">
            <Navbar/>
            <div className="flex gap-4 w-full h-full">
                <Sidebar/>
                <Outlet/>
            </div>
        </div>
    );
}

export default Layout;

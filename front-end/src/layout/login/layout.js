import {Outlet} from "react-router-dom"

function Layout() {
    return (
        <div class="flex flex-col h-full">
            <div className="flex gap-4 w-full h-full">
                <h1>Please login</h1>
                <Outlet/>
            </div>
        </div>
    );
}

export default Layout;

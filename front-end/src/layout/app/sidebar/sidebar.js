import {NavLink} from "react-router-dom";
import SidebarStyle from "./sidebar.module.css"

export default function Sidebar() {
    return (
        <div className={SidebarStyle.Sidebar}>
            <NavLink
                className={({isActive}) =>
                    isActive ? SidebarStyle.Active : ""
                }
                to="/datapoints/search">Search datapoints</NavLink>
            <NavLink
                className={({isActive}) =>
                    isActive ? SidebarStyle.Active : ""
                }
                to="/datapoints/live">Live datapoints</NavLink>
            <NavLink
                className={({isActive}) =>
                    isActive ? SidebarStyle.Active : ""
                }
                to="/devices">Devices</NavLink>
            <NavLink to="/settings">Settings</NavLink>
        </div>
    )
}
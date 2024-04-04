import {NavLink} from "react-router-dom";
import SidebarStyle from "./sidebar.module.css"
import {useSelector} from "react-redux";
import {useEffect, useState} from "react";

export default function Sidebar() {

    const application = useSelector((state) => state.application);
    const [pages, setPages] = useState({})
    const [requirements, setRequirements] = useState({
        "search": ["/api/datapoints/info/devices", "/api/datapoints/info/labels", "/api/datapoints/find"],
        "live": ["/api/datapoints/info/labels", "/ws/connect"],
        "devices": ["/api/devices/delete/*", "/api/devices/upsert", "/api/devices/get"],
        "access": ["/api/roles/upsert", "/api/roles/get", "/api/endpoints/upsert", "/api/endpoints/get", "/api/users/upsert", "/api/users/get"]
    })

    useEffect(() => {
        if (!application.user) {
            return
        }
        if (application.user.endpoints.length === 0) {
            return;
        }
        let endpoints = application.user.endpoints
        let permitted = 0
        let localPages = {
            search: false,
            live: false,
            devices: false,
            access: false,
        }
        for (const [page, reqs] of Object.entries(requirements)) {
            reqs.forEach(r => {
                let endpoint = endpoints.find(e => e.urn === r)
                if (!endpoint) {
                    return
                }
                permitted++
            })

            if (permitted === reqs.length) {
                localPages[page] = true
            }

            setPages(localPages)
            permitted = 0
        }
    }, [application.user]);

    return (
        <div className={SidebarStyle.Sidebar}>
            {
                pages.search ? <NavLink
                    className={({isActive}) =>
                        isActive ? SidebarStyle.Active : ""
                    }
                    to="/datapoints/search">Search datapoints</NavLink> : ""
            }
            {
                pages.live ?
                    <NavLink
                        className={({isActive}) =>
                            isActive ? SidebarStyle.Active : ""
                        }
                        to="/datapoints/live">Live datapoints</NavLink> : ""
            }
            {
                pages.devices ?
                    <NavLink
                        className={({isActive}) =>
                            isActive ? SidebarStyle.Active : ""
                        }
                        to="/devices">Devices</NavLink> : ""
            }
            {
                pages.access ?
                    <NavLink
                        className={({isActive}) =>
                            isActive ? SidebarStyle.Active : ""
                        }
                        to="/access-management">Access management</NavLink> : ""
            }
            <NavLink
                className={({isActive}) =>
                    isActive ? SidebarStyle.Active : ""
                }
                to="/settings">Settings</NavLink>
        </div>
    )
}
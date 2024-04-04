import {Outlet} from "react-router-dom"
import Sidebar from "./sidebar/sidebar";
import Navbar from "./navbar/navbar";
import {useSelector} from "react-redux";
import {useEffect, useState} from "react";
import Alert from "../../components/ui/alert/alert";
import {setUser} from "../../store/application/application";
import {useDispatch} from "react-redux";

function Layout() {
    const application = useSelector((state) => state.application);
    const dispatch = useDispatch()
    const [alert, setAlert] = useState(null);

    useEffect(() => {
        application.axios.get("/api/users/current-user")
            .then(e => {
                dispatch(setUser(e.data))
            })
            .catch(e => {
                setAlert({
                    type: "error",
                    description: `Cannot get current user: ${e}`,
                    link: "https://google.com"
                })
                setTimeout(() => setAlert(null), 5000)
            })
    }, [])
    return (
        <div class="flex flex-col h-full">
            <Navbar/>
            <div className="flex gap-4 w-full h-full">
                <Sidebar/>
                <Outlet/>
            </div>
            {
                alert &&
                <Alert type={alert.type} description={alert.description} link={alert.link}/>
            }
        </div>
    );
}

export default Layout;

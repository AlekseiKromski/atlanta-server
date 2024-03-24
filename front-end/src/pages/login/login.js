import {Button, Input} from "@nextui-org/react";
import {useDispatch, useSelector} from "react-redux";
import {useState} from "react";
import {addToken} from "../../store/application/application";
import {useNavigate} from "react-router-dom";
import Cookies from 'js-cookie';

export default function Login() {

    let navigate = useNavigate();
    const dispatch = useDispatch();
    const application = useSelector((state) => state.application);
    const [username, setUsername] = useState("")
    const [password, setPassword] = useState("")
    const [loader, setLoader] = useState(false)


    const login = () => {
        setLoader(true)

        application.axios.post(`/api/auth`, {
            username: username,
            password: password,
            type: "cookie"
        }, {
            withCredentials: true
        })
            .then(res => {
                dispatch(addToken(res.data));
                navigate("/")
                Cookies.set('token', res.data.token, { expires: 1, secure: true, httpOnly: true});
                setTimeout(() => setLoader(false), 1000)
            })
            .catch(e => {
                console.log(e)
                setTimeout(() => setLoader(false), 1000)
            })
    }

    return (
        <div className="">
            <Input onChange={(e) => setUsername(e.target.value)} type="text" label="Username" placeholder="Enter your username" />
            <Input onChange={(e) => setPassword(e.target.value)} type="password" label="Password" placeholder="Enter your password" />
            <Button color="primary" onClick={login} isLoading={loader}>
                Login
            </Button>
        </div>
    )
}
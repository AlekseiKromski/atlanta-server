import MainStyle from "./main.module.css"
export default function Main() {
    return (
        <div className={MainStyle.MainBody + " flex w-full"}>
            <h1>Welcome to <b>ATLANTA</b> dasboard</h1>
        </div>
    )
}
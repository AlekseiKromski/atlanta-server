import Wrapper from "../../common/wrapper/wrapper";
import UsersStyle from "./users.module.css"
import {Table as TableReact} from "@nextui-org/table";
import {
    Button, Chip,
    Dropdown,
    DropdownItem,
    DropdownMenu,
    DropdownTrigger,
    Input,
    Modal,
    ModalBody,
    ModalContent,
    ModalFooter,
    ModalHeader,
    Select,
    SelectItem,
    TableBody,
    TableCell,
    TableColumn,
    TableHeader,
    TableRow,
    useDisclosure
} from "@nextui-org/react";
import {useEffect, useState} from "react";
import {useSelector} from "react-redux";
import Alert from "../../ui/alert/alert";

export default function Users() {

    const {isOpen, onClose, onOpen, onOpenChange} = useDisclosure();
    const [user, setUser] = useState(
        {
            "username": "",
            "first_name": "",
            "second_name": "",
            "password": "",
            "email": "",
            "role": "",
            "id": "",
            "deleted_at": ""
        }
    )
    const [users, setUsers] = useState([])
    const [roles, setRoles] = useState([])
    const [updateTrigger, setUpdateTrigger] = useState(0)
    const application = useSelector((state) => state.application);
    const [alert, setAlert] = useState(null);

    const onUpsert = () => {
        if (
            user.username.length === 0 ||
            user.email.length === 0 ||
            user.first_name.length === 0 ||
            user.second_name.length === 0 ||
            user.role.length === 0
        ) {
            return
        }
        application.axios.post("/api/users/upsert", {
            "username": user.username,
            "first_name": user.first_name,
            "second_name": user.second_name,
            "password": user.password,
            "email": user.email,
            "role": user.role,
            "id": user.id.length === 0 ? null : user.id,
        })
            .then(e => {
                onClose(false)
                setUser({
                    "username": "",
                    "first_name": "",
                    "second_name": "",
                    "password": "",
                    "email": "",
                    "role": "",
                    "id": "",
                    "deleted_at": ""
                })
                setUpdateTrigger(updateTrigger + 1)
            })
            .catch(e => {
                setAlert({
                    type: "error",
                    description: `Cannot upsert user: ${e}`,
                    link: "https://google.com"
                })
                setTimeout(() => setAlert(null), 5000)
            })
    }

    const onDelete = (userId) => {
        if (!userId) {
            return
        }

        application.axios.post("/api/users/upsert", {
            "id": userId,
            "deleted_at": "ok",
        })
            .then(e => {
                onClose(false)
                setUser({
                    "username": "",
                    "first_name": "",
                    "second_name": "",
                    "password": "",
                    "email": "",
                    "role": "",
                    "id": "",
                    "deleted_at": ""
                })
                setUpdateTrigger(updateTrigger + 1)
            })
            .catch(e => {
                setAlert({
                    type: "error",
                    description: `Cannot delete user: ${e}`,
                    link: "https://google.com"
                })
                setTimeout(() => setAlert(null), 5000)
            })
    }

    const onUpdate = (userId) => {
        if (!userId) {
            return
        }

        let user = users.find(u => u.id === userId)
        if (!user) {
            return
        }

        setUser(user)
        onOpen()
    }

    useEffect(() => {
        application.axios.get("/api/users/get")
            .then(res => setUsers(res.data))
            .catch(e => {
                setAlert({
                    type: "error",
                    description: `Cannot get users: ${e}`,
                    link: "https://google.com"
                })
                setTimeout(() => setAlert(null), 5000)
            })

        application.axios.get("/api/roles/get")
            .then(res => setRoles(res.data))
            .catch(e => {
                setAlert({
                    type: "error",
                    description: `Cannot get roles: ${e}`,
                    link: "https://google.com"
                })
                setTimeout(() => setAlert(null), 5000)
            })
    }, [updateTrigger]);

    return (
        <div className={UsersStyle.UsersBody + " w-full"}>
            <Wrapper
                title="Users"
                modal={
                    {
                        title: "Users",
                        body: (
                            <p>
                                TODO: ...
                            </p>
                        )
                    }
                }
            >
                <div className={UsersStyle.UsersPanel + " flex flex-col"}>
                    <div className="">
                        <Button size="sm" color="success" variant="flat" onPress={onOpen}>Create</Button>
                        <Modal backdrop="blur" isOpen={isOpen} onClose={() => {
                            setUser({
                                "username": "",
                                "first_name": "",
                                "second_name": "",
                                "password": "",
                                "email": "",
                                "role": "",
                                "id": "",
                                "deleted_at": ""
                            })
                        }} onOpenChange={onOpenChange}>
                            <ModalContent>
                                {(onClose) => {
                                    return (
                                        (
                                            <>
                                                <ModalHeader className="flex flex-col gap-1">Create new
                                                    user</ModalHeader>
                                                <ModalBody>
                                                    <div
                                                        className={UsersStyle.UsersPanel + " flex flex-col w-full"}>

                                                        <Input
                                                            type="username"
                                                            label="Username"
                                                            className="w-full"
                                                            onChange={(e) => {
                                                                setUser({
                                                                    ...user,
                                                                    username: e.target.value,
                                                                })
                                                            }}
                                                            value={user.username}
                                                        />

                                                        <Input
                                                            type="first_name"
                                                            label="First name"
                                                            className="w-full"
                                                            onChange={(e) => {
                                                                setUser({
                                                                    ...user,
                                                                    first_name: e.target.value,
                                                                })
                                                            }}
                                                            value={user.first_name}
                                                        />

                                                        <Input
                                                            type="second_name"
                                                            label="Second name"
                                                            className="w-full"
                                                            onChange={(e) => {
                                                                setUser({
                                                                    ...user,
                                                                    second_name: e.target.value,
                                                                })
                                                            }}
                                                            value={user.second_name}
                                                        />

                                                        <Input
                                                            type="email"
                                                            label="Email"
                                                            className="w-full"
                                                            onChange={(e) => {
                                                                setUser({
                                                                    ...user,
                                                                    email: e.target.value,
                                                                })
                                                            }}
                                                            value={user.email}
                                                        />

                                                        <Input
                                                            type="password"
                                                            label="Password"
                                                            className="w-full"
                                                            onChange={(e) => {
                                                                setUser({
                                                                    ...user,
                                                                    password: e.target.value,
                                                                })
                                                            }}
                                                            value={user.password}
                                                        />

                                                        <Select
                                                            label="Role"
                                                            placeholder="Select role"
                                                            className="w-full"
                                                            onChange={(e) => {
                                                                setUser({
                                                                    ...user,
                                                                    role: e.target.value,
                                                                })
                                                            }}
                                                            selectedKeys={[user.role]}
                                                        >
                                                            {
                                                                roles && roles.length !== 0 && roles.map(role => (
                                                                    <SelectItem key={role.id}
                                                                                value={role.id}>{role.name}</SelectItem>
                                                                ))
                                                            }
                                                        </Select>
                                                    </div>
                                                </ModalBody>
                                                <ModalFooter>
                                                    <Button color="danger" variant="light" onPress={onClose}>
                                                        Close
                                                    </Button>
                                                    <Button color={user.id.length === 0 ? "success" : "warning"}
                                                            variant="light" onPress={onUpsert}>
                                                        {user.id.length === 0 ? "Create" : "Update"}
                                                    </Button>
                                                </ModalFooter>
                                            </>
                                        )
                                    )
                                }}
                            </ModalContent>
                        </Modal>
                    </div>
                    {
                        users && <TableReact aria-label="Static Actions">
                            <TableHeader>
                                <TableColumn>User ID</TableColumn>
                                <TableColumn>Username</TableColumn>
                                <TableColumn>First name</TableColumn>
                                <TableColumn>Second name</TableColumn>
                                <TableColumn>Email</TableColumn>
                                <TableColumn>Role</TableColumn>
                                <TableColumn>Created</TableColumn>
                                <TableColumn>Updated</TableColumn>
                                <TableColumn>Actions</TableColumn>
                            </TableHeader>
                            <TableBody emptyContent={"No data received"} items={users}>
                                {users.length !== 0 && roles.length !== 0 && users.map(u => (
                                    <TableRow key={u.id}>
                                        <TableCell>{u.id}</TableCell>
                                        <TableCell>{u.username}</TableCell>
                                        <TableCell>{u.first_name}</TableCell>
                                        <TableCell>{u.second_name}</TableCell>
                                        <TableCell>{u.email}</TableCell>
                                        <TableCell>
                                            <Chip color="success" variant="flat">{u.role_name}</Chip>
                                        </TableCell>
                                        <TableCell>{u.created_at}</TableCell>
                                        <TableCell>{u.updated_at}</TableCell>
                                        <TableCell>
                                            <Dropdown>
                                                <DropdownTrigger>
                                                    <Button
                                                        size="sm"
                                                        color="default" variant="flat"
                                                    >
                                                        Actions
                                                    </Button>
                                                </DropdownTrigger>
                                                <DropdownMenu selectionMode="single" onSelectionChange={(e) => {
                                                    if (e.currentKey === "delete") {
                                                        onDelete(u.id)
                                                        return
                                                    } else if (e.currentKey === "edit") {
                                                        onUpdate(u.id)
                                                        return
                                                    }
                                                }} aria-label="Static Actions">
                                                    <DropdownItem key="edit">Edit</DropdownItem>
                                                    <DropdownItem value={u.id} key="delete" className="text-danger"
                                                                  color="danger">
                                                        Delete
                                                    </DropdownItem>
                                                </DropdownMenu>
                                            </Dropdown>
                                        </TableCell>
                                    </TableRow>
                                ))}
                            </TableBody>
                        </TableReact>
                    }
                </div>
            </Wrapper>
            {
                alert &&
                <Alert type={alert.type} description={alert.description} link={alert.link}/>
            }
        </div>
    )
}
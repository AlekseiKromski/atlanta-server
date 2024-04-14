import Wrapper from "../../common/wrapper/wrapper";
import RolesStyle from "./roles.module.css"
import {Table as TableReact} from "@nextui-org/table";
import {
    Button,
    Chip,
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

export default function Roles() {

    const {isOpen, onClose, onOpen, onOpenChange} = useDisclosure();
    const [role, setRole] = useState(
        {
            "id": "",
            "name": "",
            "endpoints": []
        }
    )
    const [roles, setRoles] = useState([])
    const [endpoints, setEndpoints] = useState([])
    const [updateTrigger, setUpdateTrigger] = useState(0)
    const application = useSelector((state) => state.application);
    const [alert, setAlert] = useState(null);

    const onUpsert = () => {
        if (
            role.name.length === 0 ||
            role.endpoints.length === 0
        ) {
            return
        }
        application.axios.post("/api/roles/upsert", {
            "endpoints": role.endpoints,
            "name": role.name,
            "id": role.id.length === 0 ? null : role.id,
        })
            .then(e => {
                onClose(false)
                setRole({
                    "id": "",
                    "name": "",
                    "endpoints": []
                })
                setUpdateTrigger(updateTrigger + 1)
            })
            .catch(e => {
                setAlert({
                    type: "error",
                    description: `Cannot upsert role: ${e}`,
                    link: "https://google.com"
                })
                setTimeout(() => setAlert(null), 5000)
            })
    }

    const onDelete = (roleId) => {
        if (!roleId) {
            return
        }

        application.axios.post("/api/roles/upsert", {
            "id": roleId,
            "deleted_at": "ok"
        })
            .then(e => {
                onClose(false)
                setRole({
                    "id": "",
                    "name": "",
                    "endpoints": []
                })
                setUpdateTrigger(updateTrigger + 1)
            })
            .catch(e => {
                setAlert({
                    type: "error",
                    description: `Cannot delete role: ${e}`,
                    link: "https://google.com"
                })
                setTimeout(() => setAlert(null), 5000)
            })
    }

    const onUpdate = (roleId) => {
        if (!roleId) {
            return
        }

        let role = roles.find(r => r.id === roleId)
        if (!role) {
            return
        }

        setRole(role)
        onOpen()
    }

    useEffect(() => {
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

        application.axios.get("/api/endpoints/get")
            .then(res => setEndpoints(res.data))
            .catch(e => {
                setAlert({
                    type: "error",
                    description: `Cannot get endpoints: ${e}`,
                    link: "https://google.com"
                })
                setTimeout(() => setAlert(null), 5000)
            })
    }, [updateTrigger]);

    return (
        <div className={RolesStyle.RolesPanel + " w-full"}>
            <Wrapper
                title="Roles"
                fileName="access_management_roles.md"
                modal={
                    {
                        title: "Roles"
                    }
                }
            >
                <div className={RolesStyle.RolesPanel + " flex flex-col"}>
                    <div className="">
                        <Button size="sm" color="success" variant="flat" onPress={onOpen}>Create</Button>
                        <Modal backdrop="blur" isOpen={isOpen} onClose={() => {
                            setRole({
                                "id": "",
                                "name": "",
                                "endpoints": []
                            })
                        }} onOpenChange={onOpenChange}>
                            <ModalContent>
                                {(onClose) => {
                                    return (
                                        (
                                            <>
                                                <ModalHeader className="flex flex-col gap-1">Create new
                                                    role</ModalHeader>
                                                <ModalBody>
                                                    <div
                                                        className={RolesStyle.RolesPanel + " flex flex-col w-full"}>

                                                        <Input
                                                            type="name"
                                                            label="Name"
                                                            className="w-full"
                                                            onChange={(e) => {
                                                                setRole({
                                                                    ...role,
                                                                    name: e.target.value,
                                                                })
                                                            }}
                                                            value={role.name}
                                                        />

                                                        <Select
                                                            label="Endpoints"
                                                            placeholder="Select endpoints"
                                                            className="w-full"
                                                            onSelectionChange={(e) => {
                                                                if (e.currentKey === "all") {

                                                                    // If all labels already set, we should remove it
                                                                    if (role.endpoints.length == endpoints.length) {
                                                                        setRole({
                                                                            ...role,
                                                                            endpoints: []
                                                                        })
                                                                        return
                                                                    }

                                                                    setRole({
                                                                        ...role,
                                                                        endpoints: [...endpoints]
                                                                    })
                                                                    return
                                                                }

                                                                if (role.endpoints.find(ep => ep.id === e.currentKey)) {
                                                                    return;
                                                                }

                                                                setRole({
                                                                    ...role,
                                                                    endpoints: [...role.endpoints, endpoints.find(ep => ep.id === e.currentKey)]
                                                                })
                                                            }}
                                                        >
                                                            <SelectItem selected key="all" value="all">
                                                                All
                                                            </SelectItem>
                                                            {
                                                                endpoints && endpoints.length !== 0 && endpoints.map(e => (
                                                                    <SelectItem key={e.id}
                                                                                value={e.id}>{e.urn} - {e.description}</SelectItem>
                                                                ))
                                                            }
                                                        </Select>
                                                        <div className={RolesStyle.RolesChip + " flex flex-wrap"}>
                                                            {role.endpoints.map(e => <Chip
                                                                className={RolesStyle.RoleClick} onClick={
                                                                () => {
                                                                    setRole({
                                                                        ...role,
                                                                        endpoints: [...role.endpoints.filter(ep => ep.id !== e.id)]
                                                                    })
                                                                }
                                                            } color="success"
                                                                variant="flat">{e.urn} - <b>{e.description}</b></Chip>)}
                                                        </div>
                                                    </div>
                                                </ModalBody>
                                                <ModalFooter>
                                                    <Button color="danger" variant="light" onPress={onClose}>
                                                        Close
                                                    </Button>
                                                    <Button color={role.id.length === 0 ? "success" : "warning"}
                                                            variant="light" onPress={onUpsert}>
                                                        {role.id.length === 0 ? "Create" : "Update"}
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
                        roles && <TableReact aria-label="Static Actions">
                            <TableHeader>
                                <TableColumn>Role ID</TableColumn>
                                <TableColumn>Name</TableColumn>
                                <TableColumn>Endpoints</TableColumn>
                                <TableColumn>Created</TableColumn>
                                <TableColumn>Updated</TableColumn>
                                <TableColumn>Actions</TableColumn>
                            </TableHeader>
                            <TableBody emptyContent={"No data received"}>
                                {roles.length !== 0 && roles.map(r => (
                                    <TableRow key={r.id}>
                                        <TableCell>{r.id}</TableCell>
                                        <TableCell>{r.name}</TableCell>
                                        <TableCell>
                                            <div className={RolesStyle.RolesChip + " flex flex-wrap"}>
                                                {r.endpoints.map(e => <Chip color="success"
                                                                            variant="flat">{e.urn} - <b>{e.description}</b></Chip>)}
                                            </div>
                                        </TableCell>
                                        <TableCell>{r.created_at}</TableCell>
                                        <TableCell>{r.updated_at}</TableCell>
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
                                                        onDelete(r.id)
                                                        return
                                                    } else if (e.currentKey === "edit") {
                                                        onUpdate(r.id)
                                                        return
                                                    }
                                                }} aria-label="Static Actions">
                                                    <DropdownItem key="edit">Edit</DropdownItem>
                                                    <DropdownItem value={r.id} key="delete" className="text-danger"
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
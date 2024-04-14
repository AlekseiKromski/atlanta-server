import Wrapper from "../../components/common/wrapper/wrapper";
import DevicesStyle from "./devices.module.css"
import {Table as TableReact} from "@nextui-org/table";
import {
    Button,
    Chip,
    Dropdown,
    DropdownItem,
    DropdownMenu,
    DropdownTrigger,
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
    Textarea,
    useDisclosure
} from "@nextui-org/react";
import {useEffect, useState} from "react";
import {useSelector} from "react-redux";
import Alert from "../../components/ui/alert/alert";

export default function Devices() {

    const {isOpen, onClose, onOpen, onOpenChange} = useDisclosure();
    const [device, setDevice] = useState({
        description: "",
        status: "",
        id: ""
    })
    const [devices, setDevices] = useState([])
    const [updateTrigger, setUpdateTrigger] = useState(0)
    const application = useSelector((state) => state.application);
    const [alert, setAlert] = useState(null);

    const onCreate = () => {
        if (device.description.length === 0 || device.status.length === 0) {
            return
        }
        application.axios.post("/api/devices/upsert", {
            "description": device.description,
            "status": device.status === "true",
            "id": device.id.length !== 0 ? device.id : null
        })
            .then(e => {
                onClose(false)
                setDevice({
                    description: "",
                    status: "",
                    id: ""
                })
                setUpdateTrigger(updateTrigger + 1)
            })
            .catch(e => {
                setAlert({
                    type: "error",
                    description: `Cannot upsert device: ${e}`,
                    link: "https://google.com"
                })
                setTimeout(() => setAlert(null), 5000)
            })
    }

    const onDelete = (deviceId) => {
        if (!deviceId) {
            return
        }

        application.axios.get("/api/devices/delete/" + deviceId)
            .then(e => {
                setUpdateTrigger(updateTrigger + 1)
            })
            .catch(e => {
                setAlert({
                    type: "error",
                    description: `Cannot delete device: ${e}`,
                    link: "https://google.com"
                })
                setTimeout(() => setAlert(null), 5000)
            })
    }

    const onUpdate = (deviceId) => {
        if (!deviceId) {
            return
        }

        let device = devices.find(d => d.id === deviceId)
        if (!device) {
            return
        }

        setDevice(device)
        onOpen()
    }

    useEffect(() => {
        application.axios.get("/api/devices/get")
            .then(res => setDevices(res.data))
            .catch(e => {
                setAlert({
                    type: "error",
                    description: `Cannot get devices: ${e}`,
                    link: "https://google.com"
                })
                setTimeout(() => setAlert(null), 5000)
            })
    }, [updateTrigger]);

    return (
        <div className={DevicesStyle.DevicesBody + " w-full"}>
            <Wrapper
                title="Devices"
                fileName="devices.md"
                modal={
                    {
                        title: "Devices"
                    }
                }
            >
                <div className={DevicesStyle.DevicesPanel + " flex flex-col"}>
                    <div className="">
                        <Button size="sm" color="success" variant="flat" onPress={onOpen}>Create</Button>
                        <Modal backdrop="blur" isOpen={isOpen} onClose={() => {
                            setDevice({
                                description: "",
                                status: "",
                                id: ""
                            })
                        }} onOpenChange={onOpenChange}>
                            <ModalContent>
                                {(onClose) => {
                                    return (
                                        (
                                            <>
                                                <ModalHeader className="flex flex-col gap-1">Create new
                                                    device</ModalHeader>
                                                <ModalBody>
                                                    <div
                                                        className={DevicesStyle.DevicesPanel + " flex flex-col w-full"}>
                                                        <Select
                                                            label="Status"
                                                            placeholder="Select device status"
                                                            className="w-full"
                                                            onChange={(e) => {
                                                                setDevice({
                                                                    ...device,
                                                                    status: e.target.value,
                                                                })
                                                            }}
                                                            selectedKeys={[device.status.toString()]}
                                                        >
                                                            <SelectItem key="true">Enabled</SelectItem>
                                                            <SelectItem key="false">Disabled</SelectItem>
                                                        </Select>
                                                        <Textarea
                                                            label="Description"
                                                            placeholder="Enter your description"
                                                            className="w-full"
                                                            value={device.description}
                                                            onChange={(e) => {
                                                                setDevice({
                                                                    ...device,
                                                                    description: e.target.value,
                                                                })
                                                            }}
                                                        />
                                                    </div>
                                                </ModalBody>
                                                <ModalFooter>
                                                    <Button color="danger" variant="light" onPress={onClose}>
                                                        Close
                                                    </Button>
                                                    <Button color={device.id.length === 0 ? "success" : "warning"}
                                                            variant="light" onPress={onCreate}>
                                                        {device.id.length === 0 ? "Create" : "Update"}
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
                        devices && <TableReact aria-label="Static Actions">
                            <TableHeader>
                                <TableColumn>Device ID</TableColumn>
                                <TableColumn>Description</TableColumn>
                                <TableColumn>Status</TableColumn>
                                <TableColumn>Created</TableColumn>
                                <TableColumn>Updated</TableColumn>
                                <TableColumn>Actions</TableColumn>
                            </TableHeader>
                            <TableBody emptyContent={"No data received"} items={devices}>
                                {devices.length !== 0 && devices.map(d => (
                                    <TableRow key={d.id}>
                                        <TableCell>{d.id}</TableCell>
                                        <TableCell>{d.description}</TableCell>
                                        <TableCell>
                                            {
                                                d.status ?
                                                    <Chip color="success" variant="flat">Enabled</Chip>
                                                    : <Chip color="danger" variant="flat">Disabled</Chip>
                                            }
                                        </TableCell>
                                        <TableCell>{d.created_at}</TableCell>
                                        <TableCell>{d.updated_at}</TableCell>
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
                                                        onDelete(d.id)
                                                        return
                                                    } else if (e.currentKey === "edit") {
                                                        onUpdate(d.id)
                                                        return
                                                    }
                                                }} aria-label="Static Actions">
                                                    <DropdownItem key="edit">Edit</DropdownItem>
                                                    <DropdownItem value={d.id} key="delete" className="text-danger"
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
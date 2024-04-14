import Wrapper from "../../common/wrapper/wrapper";
import EndpointStyle from "./endpoints.module.css"
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
import Alert from "../../ui/alert/alert";

export default function Endpoints() {

    const {isOpen, onClose, onOpen, onOpenChange} = useDisclosure();
    const [endpoint, setEndpoint] = useState(
        {
            "id": "",
            "urn": "",
            "description": "",
            "deleted_at": ""
        }
    )
    const [endpoints, setEndpoints] = useState([])
    const [updateTrigger, setUpdateTrigger] = useState(0)
    const application = useSelector((state) => state.application);
    const [alert, setAlert] = useState(null);

    const onUpsert = () => {
        if (
            endpoint.urn.length === 0 ||
            endpoint.description.length === 0
        ) {
            return
        }
        application.axios.post("/api/endpoints/upsert",
            {
                "id": endpoint.id.length === 0 ? null : endpoint.id,
                "urn": endpoint.urn,
                "description": endpoint.description,
            })
            .then(e => {
                onClose(false)
                setEndpoint({
                    "id": "",
                    "urn": "",
                    "description": "",
                    "deleted_at": ""
                })
                setUpdateTrigger(updateTrigger + 1)
            })
            .catch(e => {
                setAlert({
                    type: "error",
                    description: `Cannot upsert endpoint: ${e}`,
                    link: "https://google.com"
                })
                setTimeout(() => setAlert(null), 5000)
            })
    }

    const onDelete = (endpointId) => {
        if (!endpointId) {
            return
        }

        application.axios.post("/api/endpoints/upsert", {
            "id": endpointId,
            "deleted_at": "ok",
        })
            .then(e => {
                onClose(false)
                setEndpoint({
                    "id": "",
                    "urn": "",
                    "description": "",
                    "deleted_at": ""
                })
                setUpdateTrigger(updateTrigger + 1)
            })
            .catch(e => {
                setAlert({
                    type: "error",
                    description: `Cannot delete endpoint: ${e}`,
                    link: "https://google.com"
                })
                setTimeout(() => setAlert(null), 5000)
            })
    }

    const onUpdate = (endpointId) => {
        if (!endpointId) {
            return
        }

        let endpoint = endpoints.find(e => e.id === endpointId)
        if (!endpoint) {
            return
        }

        setEndpoint(endpoint)
        onOpen()
    }

    useEffect(() => {
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
        <div className={EndpointStyle.EndpointsBody + " w-full"}>
            <Wrapper
                title="Endpoints"
                fileName="access_management_endpoints.md"
                modal={
                    {
                        title: "Endpoints"
                    }
                }
            >
                <div className={EndpointStyle.EndpointsPanel + " flex flex-col"}>
                    <div className="">
                        <Button size="sm" color="success" variant="flat" onPress={onOpen}>Create</Button>
                        <Modal backdrop="blur" isOpen={isOpen} onClose={() => {
                            setEndpoint({
                                "id": "",
                                "urn": "",
                                "description": "",
                                "deleted_at": ""
                            })
                        }} onOpenChange={onOpenChange}>
                            <ModalContent>
                                {(onClose) => {
                                    return (
                                        (
                                            <>
                                                <ModalHeader className="flex flex-col gap-1">Create new
                                                    endpoint</ModalHeader>
                                                <ModalBody>
                                                    <div
                                                        className={EndpointStyle.EndpointsPanel + " flex flex-col w-full"}>
                                                        <Input
                                                            type="urn"
                                                            label="Urn"
                                                            className="w-full"
                                                            onChange={(e) => {
                                                                setEndpoint({
                                                                    ...endpoint,
                                                                    urn: e.target.value,
                                                                })
                                                            }}
                                                            value={endpoint.urn}
                                                        />

                                                        <Textarea
                                                            label="Description"
                                                            placeholder="Enter your description"
                                                            className="w-full"
                                                            value={endpoint.description}
                                                            onChange={(e) => {
                                                                setEndpoint({
                                                                    ...endpoint,
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
                                                    <Button color={endpoint.id.length === 0 ? "success" : "warning"}
                                                            variant="light" onPress={onUpsert}>
                                                        {endpoint.id.length === 0 ? "Create" : "Update"}
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
                        endpoints && <TableReact aria-label="Static Actions">
                            <TableHeader>
                                <TableColumn>Endpoint ID</TableColumn>
                                <TableColumn>Urn</TableColumn>
                                <TableColumn>Description</TableColumn>
                                <TableColumn>Created</TableColumn>
                                <TableColumn>Updated</TableColumn>
                                <TableColumn>Actions</TableColumn>
                            </TableHeader>
                            <TableBody emptyContent={"No data received"} items={endpoints}>
                                {endpoints.length !== 0 && endpoints.map(endpoint => (
                                    <TableRow key={endpoint.id}>
                                        <TableCell>{endpoint.id}</TableCell>
                                        <TableCell>{endpoint.urn}</TableCell>
                                        <TableCell>
                                            <Chip color="success" variant="flat">{endpoint.description}</Chip>
                                        </TableCell>
                                        <TableCell>{endpoint.created_at}</TableCell>
                                        <TableCell>{endpoint.updated_at}</TableCell>
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
                                                        onDelete(endpoint.id)
                                                        return
                                                    } else if (e.currentKey === "edit") {
                                                        onUpdate(endpoint.id)
                                                        return
                                                    }
                                                }} aria-label="Static Actions">
                                                    <DropdownItem key="edit">Edit</DropdownItem>
                                                    <DropdownItem value={endpoint.id} key="delete" className="text-danger"
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
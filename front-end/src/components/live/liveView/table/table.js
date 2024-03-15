import {
    Button,
    Modal, ModalBody,
    ModalContent, ModalFooter,
    ModalHeader, Select, SelectItem, Tab,
    Table as TableReact,
    TableBody,
    TableCell,
    TableColumn,
    TableHeader,
    TableRow, useDisclosure
} from "@nextui-org/react";
import TableStyle from "./table.module.css";
export default function Table({wrapper, device, datapoints}) {

    const {isOpen, onOpen, onOpenChange} = useDisclosure();

    return (
        <div className={TableStyle.Main + " w-full flex flex-col" + wrapper ? TableStyle.Wrapper : ""}>
            <div className={TableStyle.Header + " flex justify-between items-center rounded-md"}>
                <h1 className="font-bold">Last datapoints table</h1>
                <Button size="sm" color="primary" variant="light" onPress={onOpen}>Help</Button>
                <Modal backdrop="blur" isOpen={isOpen} onOpenChange={onOpenChange}>
                    <ModalContent>
                        {(onClose) => (
                            <>
                                <ModalHeader className="flex flex-col gap-1">Search instruction</ModalHeader>
                                <ModalBody>
                                    <p>
                                        Lorem ipsum dolor sit amet, consectetur adipiscing elit.
                                        Nullam pulvinar risus non risus hendrerit venenatis.
                                        Pellentesque sit amet hendrerit risus, sed porttitor quam.
                                    </p>
                                    <p>
                                        Lorem ipsum dolor sit amet, consectetur adipiscing elit.
                                        Nullam pulvinar risus non risus hendrerit venenatis.
                                        Pellentesque sit amet hendrerit risus, sed porttitor quam.
                                    </p>
                                    <p>
                                        Magna exercitation reprehenderit magna aute tempor cupidatat consequat elit
                                        dolor adipisicing. Mollit dolor eiusmod sunt ex incididunt cillum quis.
                                        Velit duis sit officia eiusmod Lorem aliqua enim laboris do dolor eiusmod.
                                        Et mollit incididunt nisi consectetur esse laborum eiusmod pariatur
                                        proident Lorem eiusmod et. Culpa deserunt nostrud ad veniam.
                                    </p>
                                </ModalBody>
                                <ModalFooter>
                                    <Button color="danger" variant="light" onPress={onClose}>
                                        Close
                                    </Button>
                                </ModalFooter>
                            </>
                        )}
                    </ModalContent>
                </Modal>
            </div>
            <div className={TableStyle.Body}>
                {
                    datapoints &&  <TableReact className={TableStyle.Table} aria-label="Example static collection table">
                        <TableHeader>
                            <TableColumn>Device ID</TableColumn>
                            <TableColumn>Label</TableColumn>
                            <TableColumn>Value</TableColumn>
                            <TableColumn>Unit</TableColumn>
                            <TableColumn>Flags</TableColumn>
                            <TableColumn>Measurement time</TableColumn>
                            <TableColumn>Created</TableColumn>
                        </TableHeader>
                        <TableBody emptyContent={"No data received"} items={datapoints}>
                            {datapoints.length !== 0 && datapoints.filter(dp => dp.device_id === device).map(dp => (
                                <TableRow key={dp.id}>
                                    <TableCell>{dp.device_id}</TableCell>
                                    <TableCell>{dp.label}</TableCell>
                                    <TableCell>{dp.value}</TableCell>
                                    <TableCell>{dp.unit}</TableCell>
                                    <TableCell>{dp.flags}</TableCell>
                                    <TableCell>{dp.measurement_time}</TableCell>
                                    <TableCell>{dp.created_at}</TableCell>
                                </TableRow>
                            ))}
                        </TableBody>
                    </TableReact>
                }
            </div>
        </div>
    )
}
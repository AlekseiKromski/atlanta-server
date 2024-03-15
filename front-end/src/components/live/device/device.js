import {
    Button,
    Modal, ModalBody,
    ModalContent, ModalFooter,
    ModalHeader, Select, SelectItem,
    useDisclosure
} from "@nextui-org/react";
import DeviceStyle from "./device.module.css";

export default function Device({device, setDevice, devices}) {

    const {isOpen, onOpen, onOpenChange} = useDisclosure();

    return (
        <div className={DeviceStyle.Main + " w-full flex flex-col"}>
            <div className={DeviceStyle.Header + " flex justify-between items-center rounded-md"}>
                <h1 className="font-bold">Select device</h1>
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
            <div className={DeviceStyle.Body}>
                <Select
                    label="Select device"
                    className="max-w-full"
                    onChange={e => setDevice(e.target.value)}
                    value={device}
                >
                    {
                        devices
                            .map(device => (
                                <SelectItem key={device} value={device}>
                                    {device}
                                </SelectItem>
                            ))
                    }
                </Select>
            </div>
        </div>
    )
}
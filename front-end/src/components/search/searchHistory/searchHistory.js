import SearchHistoryStyle from "./searchHistory.module.css"
import SearchBoxStyle from "../searchBox/searchBox.module.css";
import {
    Button,
    Modal,
    ModalBody,
    ModalContent,
    ModalFooter,
    ModalHeader,
    Select, SelectItem,
    useDisclosure
} from "@nextui-org/react";

export default function SearchHistory() {
    const {isOpen, onOpen, onOpenChange} = useDisclosure();

    return (
        <div className={SearchHistoryStyle.SearchHistory}>
            <div className={SearchBoxStyle.Header + " flex justify-between items-center rounded-md"}>
                <h1 className="font-bold">History</h1>
                <Button size="sm" color="primary" variant="light" onPress={onOpen}>Help</Button>
                <Modal backdrop="blur" isOpen={isOpen} onOpenChange={onOpenChange}>
                    <ModalContent>
                        {(onClose) => (
                            <>
                                <ModalHeader className="flex flex-col gap-1">History instruction</ModalHeader>
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
            <div class="flex flex-col justify-between h-full">
                <div>
                    <Select
                        label="Select record"
                        className="max-w-xs"
                    >
                        <SelectItem key="map" value="map">
                            Map
                        </SelectItem>
                    </Select>
                </div>
            </div>
        </div>
    )
}
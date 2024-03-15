import WrapperStyle from "./wrapper.module.css"
import {
    Button,
    Modal,
    ModalBody,
    ModalContent,
    ModalFooter,
    ModalHeader,
    useDisclosure
} from "@nextui-org/react";

export default function Wrapper({children, isDark, width, height, title, modal}) {

    const {isOpen, onOpen, onOpenChange} = useDisclosure();

    return (
        <div style={{
            width: width ? width : "100%",
            height: height ? height : "auto",
        }} className={ isDark ? WrapperStyle.Wrapper : WrapperStyle.Main + " w-full flex flex-col"}>
            <div className={WrapperStyle.Header + " flex justify-between items-center rounded-md"}>
                <h1 className="font-bold">{title}</h1>
                <Button size="sm" color="primary" variant="light" onPress={onOpen}>Help</Button>
                <Modal backdrop="blur" isOpen={isOpen} onOpenChange={onOpenChange}>
                    <ModalContent>
                        {(onClose) => (
                            <>
                                <ModalHeader className="flex flex-col gap-1">{modal.title}</ModalHeader>
                                <ModalBody>
                                    {modal.body}
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
            <div className={WrapperStyle.Body}>
                {children}
            </div>
        </div>

    )
}
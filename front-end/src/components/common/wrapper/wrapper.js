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
import Markdown from 'react-markdown'
import {useEffect, useState} from "react";

export default function Wrapper({children, isDark, width, height, title, modal, fileName}) {

    const {isOpen, onOpen, onOpenChange} = useDisclosure();
    const [text, setText] = useState("")

    useEffect(() => {
        if (fileName) {
            fetch(`/docs/${fileName}`)
                .then(r => {
                    r.text().then(
                        t => {
                            setText(t)
                        }
                    )
                })
        }else {
            setText("Cannot get documentation")
        }
    }, []);

    return (
        <div style={{
            width: width ? width : "100%",
            height: height ? height : "auto",
        }} className={ isDark ? WrapperStyle.Wrapper : WrapperStyle.Main + " w-full flex flex-col"}>
            <div className={WrapperStyle.Header + " flex justify-between items-center rounded-md"}>
                <h1 className="font-bold">{title}</h1>
                <Button size="sm" color="primary" variant="light" onPress={onOpen}>Help</Button>
                <Modal size={"5xl"} backdrop="blur" isOpen={isOpen} onOpenChange={onOpenChange}>
                    <ModalContent>
                        {(onClose) => (
                            <>
                                <ModalHeader className="flex flex-col gap-1">{modal.title}</ModalHeader>
                                <ModalBody>
                                    <Markdown className="Markdown">
                                        {text}
                                    </Markdown>
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
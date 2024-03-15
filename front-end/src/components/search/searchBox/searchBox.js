import SearchBoxStyle from "./searchBox.module.css"
import {
    Button, Checkbox, Chip,
    Input,
    Modal,
    ModalBody,
    ModalContent,
    ModalFooter,
    ModalHeader, Select, SelectItem,
    useDisclosure
} from "@nextui-org/react";
import {useState} from "react";
import {useSelector} from "react-redux"
import moment from "moment"

export default function SearchBox({labels, callback}) {

    const application = useSelector((state) => state.application);

    const [startDate, setStartDate] = useState("")
    const [endDate, setEndDate] = useState("")
    const [selectedDatapointTypes, setSelectedDatapointTypes] = useState([])
    const [showType, setShowType] = useState("")
    const [ignored, setIgnored] = useState(false)

    const [loader, setLoader] = useState(false)

    const {isOpen, onOpen, onOpenChange} = useDisclosure();

    const find = () => {
        setLoader(true)

        let start = startDate !== "" ? moment(startDate).format().split("+")[0] + "Z" : null
        let end = endDate !== "" ? moment(endDate).format().split("+")[0] + "Z" : null
        application.axios.post(`/api/datapoints/find`, {
            start: start,
            end: end,
            select: selectedDatapointTypes,
            ignored: ignored
        })
            .then(res => {
                callback({
                    type: showType,
                    data: res.data
                })
                setTimeout(() => setLoader(false), 1000)
            })
            .catch(e => {
                console.log(e)
                setTimeout(() => setLoader(false), 1000)
            })
    }

    return (
        <div className={SearchBoxStyle.SearchBox + " w-full flex flex-col"}>
            <div className={SearchBoxStyle.Header + " flex justify-between items-center rounded-md"}>
                <h1 className="font-bold">Search</h1>
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
            <div className={SearchBoxStyle.Body}>
                <div class="flex w-full gap-1.5">
                    <Input
                        labelPlacement="inside"
                        type="datetime-local"
                        placeholder="none"
                        label="Start"
                        className="max-w-full"
                        onChange={e => setStartDate(e.target.value)}
                        value={startDate}
                    />
                    <Input
                        type="datetime-local"
                        label="End"
                        className="max-w-full"
                        placeholder="none"
                        onChange={e => setEndDate(e.target.value)}
                        value={endDate}
                    />
                </div>
                <div className={ SearchBoxStyle.SelectDatapointType + " flex w-full gap-1.5 flex-col"}>
                    <Select
                        label="Select datapoint type"
                        selectionMode="multiple"
                        className="max-w-max"
                        selectedKeys={selectedDatapointTypes}
                        onChange={e => {
                            if (e.target.value.search("all") !== -1 ) {

                                // If all labels already set, we should remove it
                                if (selectedDatapointTypes.length == labels.length) {
                                    setSelectedDatapointTypes([])
                                    return
                                }

                                setSelectedDatapointTypes(labels)
                                return
                            }

                            let values = e.target.value.split(",")
                            setSelectedDatapointTypes(values)
                        }}
                    >
                        <SelectItem selected  key="all" value="all">
                            All
                        </SelectItem>
                        {
                            labels && labels.map(label => (
                                <SelectItem selected  key={label} value={label}>
                                    {label}
                                </SelectItem>
                            ))
                        }
                    </Select>
                </div>
                <div className={ SearchBoxStyle.SelectDatapointType + " flex w-full gap-1.5"}>
                    <Select
                        label="Select show type"
                        className="max-w-xs"
                        onChange={e => setShowType(e.target.value)}
                        selectedKeys={[showType]}
                    >
                        <SelectItem key="map" value="map">
                            Map
                        </SelectItem>
                        <SelectItem key="chart" value="chart">
                            Chart
                        </SelectItem>
                    </Select>
                </div>
                <div className={ SearchBoxStyle.SelectDatapointType + " flex w-full gap-1.5"}>
                    <Checkbox isSelected={ignored} onChange={e => setIgnored(e.target.checked)} color="warning">With ignored</Checkbox>
                </div>
            </div>
            <div class="flex justify-end gap-1">

                <Button color="default" variant="flat" onClick={e => {
                    setStartDate("")
                    setEndDate("")
                    setSelectedDatapointTypes([])
                    setShowType("")
                    setIgnored(false)
                    callback({
                        type: "",
                        data: []
                    })
                }}>
                    Clear
                </Button>

                <Button isDisabled={
                    startDate === "" || endDate == "" || selectedDatapointTypes.length === 0 || showType === ""
                } isLoading={loader} color="success" variant="flat" onClick={find}>
                    Save query
                </Button>

                <Button isDisabled={
                    startDate === "" || endDate == "" || selectedDatapointTypes.length === 0 || showType === ""
                } isLoading={loader} color="secondary" onClick={find}>
                    Find
                </Button>
            </div>
        </div>
    )
}
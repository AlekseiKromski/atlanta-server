import {
    Button,
    Modal,
    ModalBody,
    ModalContent,
    ModalFooter,
    ModalHeader,
    Tab,
    Tabs, useDisclosure
} from "@nextui-org/react";
import Map from "../../search/searchChartView/map/map";
import Chart from "../../search/searchChartView/chart/chart";
import Table from "./table/table";
import LiveViewStyle from "./liveView.module.css"

export default function LiveView({datapoints, device, labels}) {

    const {isOpen, onOpen, onOpenChange} = useDisclosure();

    const prepareDatapoints = () => {
        let map = {}
        datapoints.forEach(dp => {
            if (!map[dp.label]) {
                map[dp.label] = [dp]
                return
            }

            map[dp.label] = [...map[dp.label], dp]
        })

        return map
    }

    return (
        <div className={LiveViewStyle.Main + " w-full flex flex-col"}>
            <div className={LiveViewStyle.Header + " flex justify-between items-center rounded-md"}>
                <h1 className="font-bold">Views</h1>
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
            <div className={LiveViewStyle.Body}>
                <div className="flex flex-col">
                    <Tabs key="primary" color="primary" aria-label="Show type" radius="full">
                        <Tab key="map" title="Map">
                            <Map wrapper={true} labels={labels} datapoints={
                                {
                                    data: {
                                        datapoints: prepareDatapoints(),
                                        labels: datapoints.map(dp => dp.measurement_time).filter((value, index, array) => {
                                            return array.indexOf(value) === index;
                                        }).reverse()
                                    }
                                }
                            }/>
                        </Tab>
                        <Tab key="chart" title="Chart">
                            <Chart wrapper={true} labels={labels} datapoints={{
                                data: {
                                    datapoints: prepareDatapoints(),
                                    labels: datapoints.map(dp => dp.measurement_time).filter((value, index, array) => {
                                        return array.indexOf(value) === index;
                                    }).reverse()
                                }
                            }}/>
                        </Tab>
                        <Tab key="table" title="Table">
                            <Table device={device} datapoints={datapoints}/>
                        </Tab>
                    </Tabs>
                </div>
            </div>
        </div>
    )
}
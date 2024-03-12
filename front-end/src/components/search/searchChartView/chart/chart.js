import {useEffect, useState} from "react";
import {
    Chart as ChartJS,
    CategoryScale,
    LinearScale,
    PointElement,
    LineElement,
    Title,
    Tooltip,
    Legend,
} from 'chart.js';
import { Line } from 'react-chartjs-2';
import ChartStyle from "./charts.module.css"
import MapStyle from "../map/map.module.css";
import {
    Button,
    Modal,
    ModalBody,
    ModalContent,
    ModalFooter,
    ModalHeader,
    Popover, PopoverContent, PopoverTrigger,
    useDisclosure
} from "@nextui-org/react";

ChartJS.register(
    CategoryScale,
    LinearScale,
    PointElement,
    LineElement,
    Title,
    Tooltip,
    Legend
);

function random(min, max) {
    return Math.random() * (max - min) + min;
}

export default function Chart({labels, datapoints}) {
    const {isOpen, onOpen, onOpenChange} = useDisclosure();

    datapoints = datapoints.data

    const options = {
        responsive: true,
        plugins: {
            legend: {
                position: 'top',
            },
            title: {
                display: true,
                text: 'Data',
            },
        },
    };

    let datasets = []
    for (const [key, value] of Object.entries(datapoints.datapoints)) {
        let red = random(0,255)
        let green = random(0,255)
        let blue = random(0,255)
        datasets.unshift({
            label: key,
            data: value.map(item => item.value),
            borderColor: `rgb(${red}, ${green}, ${blue})`,
            backgroundColor: `rgb(${red}, ${green}, ${blue}, 0.5)`,
        })
    }

    let datalist = {
        labels: datapoints.labels,
        datasets: datasets
    };

    let lastLength = null
    let result = true
    labels.forEach(label => {
        if (!datapoints.datapoints[label]) {
            return;
        }
        let currentLength = datapoints.datapoints[label].length
        if (lastLength == null) {
            lastLength = currentLength
            return
        }

        if (currentLength !== lastLength ){
            result = false
        }
    })

    return (
        <div>
            <div className={ChartStyle.Header + " flex justify-between items-center rounded-md"}>
                <h1 className="font-bold">Chart view</h1>
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
            <div className={ChartStyle.Chart + " flex flex-col"}>
                <div>
                    <Popover placement="right">
                        <PopoverTrigger>
                            {
                                <Button color={result ? "success" : "danger"}>Show state {result ? "✅" : "⛔"}</Button>
                            }
                        </PopoverTrigger>
                        <PopoverContent>
                            <b>Count of records</b>
                            <ul>
                                {
                                    labels && labels.length !== 0 &&
                                    labels.map(label => {
                                        if (datapoints.datapoints[label]) {
                                            return <li>Count {label} records: <b>{datapoints.datapoints[label].length}</b></li>
                                        }
                                    })
                                }
                            </ul>
                        </PopoverContent>
                    </Popover>
                </div>
                <Line options={options} data={datalist} />
            </div>
        </div>
    );
}
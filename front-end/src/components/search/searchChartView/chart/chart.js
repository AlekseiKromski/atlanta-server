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
import {
    Button,
    Popover, PopoverContent, PopoverTrigger,
} from "@nextui-org/react";
import Wrapper from "../../../common/wrapper/wrapper";

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

export default function Chart({isDark, labels, datapoints}) {

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

        if (currentLength !== lastLength){
            result = false
        }
    })

    return (
        <Wrapper
            title="Chart"
            fileName="live_server_chart.md"
            modal={{
                title: "Chart help"
            }}
            isDark={isDark}
        >
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
        </Wrapper>
    );
}
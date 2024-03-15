import {
    Tab,
    Tabs
} from "@nextui-org/react";
import Map from "../../search/searchChartView/map/map";
import Chart from "../../search/searchChartView/chart/chart";
import Table from "./table/table";
import Wrapper from "../../common/wrapper/wrapper";

export default function LiveView({datapoints, device, labels}) {

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
        <Wrapper
            title="Live view"
            modal={{
                title: "Live view",
                body: (
                    <p>
                        TODO: ...
                    </p>
                )
            }}
        >
            <div className="flex flex-col">
                <Tabs key="primary" color="primary" aria-label="Show type" radius="full">
                    <Tab key="map" title="Map">
                        <Map isDark={true} labels={labels} datapoints={
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
                        <Chart isDark={true} labels={labels} datapoints={{
                            data: {
                                datapoints: prepareDatapoints(),
                                labels: datapoints.map(dp => dp.measurement_time).filter((value, index, array) => {
                                    return array.indexOf(value) === index;
                                }).reverse()
                            }
                        }}/>
                    </Tab>
                    <Tab key="table" title="Table">
                        <Table isDark={true} device={device} datapoints={datapoints}/>
                    </Tab>
                </Tabs>
            </div>
        </Wrapper>
    )
}
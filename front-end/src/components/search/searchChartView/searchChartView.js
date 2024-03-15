import SearchChartViewStyle from "./searchChartView.module.css"
import Chart from "./chart/chart"
import Map from "./map/map"
import Wrapper from "../../common/wrapper/wrapper";

export default function SearchChartView({labels, datapoints}){
    return (
        <div className={SearchChartViewStyle.SearchChartView}>
            {
                datapoints.type === 'chart' &&
                <Chart labels={labels} datapoints={datapoints}/>
            }
            {
                datapoints.type === 'map' &&
                <Map labels={labels} datapoints={datapoints}/>
            }
        </div>
    )
}
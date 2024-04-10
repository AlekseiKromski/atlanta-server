import SearchStyle from "./search.module.css"
import SearchBox from "../../components/search/searchBox/searchBox";
import SearchHistory from "../../components/search/searchHistory/searchHistory";
import SearchChartView from "../../components/search/searchChartView/searchChartView";
import {useEffect, useState} from "react";
import {useSelector} from "react-redux"

export default function Search() {

    const application = useSelector((state) => state.application);

    const [datapoints, setDatapoints] = useState({
        type: "",
        data: []
    })
    const [labels, setLabels] = useState([])
    const [devices, setDevices] = useState([])
    const [history, setHistory] = useState({value:[]})
    const [searchOptions, setSearchOptions] = useState(null)

    const searchResult = (data) => setDatapoints(data)

    useEffect(() => {
        application.axios.get("/api/datapoints/info/labels")
            .then(res => setLabels(res.data))
            .catch(e => console.log(e))

        application.axios.get("/api/datapoints/info/devices")
            .then(res => setDevices(res.data))
            .catch(e => console.log(e))
        application.axios.get("/api/store/get/history")
            .then(res => {
                if (res.data.length === 0) {
                    return
                }

                let history = res.data[0]
                history.value = JSON.parse(history.value)
                setHistory(history)
            })
            .catch(e => {})
    }, []);

    const setSearchParameters = (historyTimestamp) => {
        setSearchOptions(history.value.find(r => r.name === historyTimestamp).value)
    }

    return (
        <div className={SearchStyle.SearchBody + " w-full flex flex-col"}>
            <div className={SearchStyle.Wrapper + " flex justify-between w-full"}>
                <SearchBox
                    history={history}
                    setHistory={setHistory}
                    labels={labels}
                    devices={devices}
                    callback={searchResult}
                    searchOptions={searchOptions}
                />
                <SearchHistory history={history} setHistory={setHistory} setSearchParameters={setSearchParameters} />
            </div>
            {
                datapoints.data.length !== 0 && <SearchChartView labels={labels} datapoints={datapoints}/>
            }
        </div>
    )
}
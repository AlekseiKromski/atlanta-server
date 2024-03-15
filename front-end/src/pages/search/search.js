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

    const searchResult = (data) => setDatapoints(data)

    useEffect(() => {
        application.axios.get("/api/datapoints/info/labels")
            .then(res => setLabels(res.data))
            .catch(e => console.log(e))
    }, []);

    return (
        <div className={SearchStyle.SearchBody + " w-full flex flex-col"}>
            <div className={SearchStyle.Wrapper + " flex justify-between w-full"}>
                <SearchBox labels={labels} callback={searchResult}/>
                <SearchHistory/>
            </div>
            {
                datapoints.data.length !== 0 && <SearchChartView labels={labels} datapoints={datapoints}/>
            }
        </div>
    )
}
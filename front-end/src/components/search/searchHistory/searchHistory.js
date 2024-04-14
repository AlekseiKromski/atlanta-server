import {Button, Select, SelectItem,} from "@nextui-org/react";
import Wrapper from "../../common/wrapper/wrapper";
import StyleHistory from "./searchHistory.module.css"
import {useSelector} from "react-redux";
import {useEffect, useState} from "react";
export default function SearchHistory({history, setHistory, setSearchParameters}) {
    const application = useSelector((state) => state.application);
    const [loader, setLoader] = useState(false)

    const remove = () => {
        let payload = {
            key: "history",
            value: JSON.stringify([])
        }

        application.axios.post(`/api/store/upsert`, payload)
            .then(res => {
                setHistory(
                    {
                        key: "history",
                        value: []
                    }
                )
                setTimeout(() => setLoader(false), 1000)
            })
            .catch(e => {
                setTimeout(() => setLoader(false), 1000)
            })
    }

    return (
        <Wrapper width="30%" title="History" fileName={"history.md"} modal={{
            title: "History help"
        }}>
            <div className="flex flex-col h-full">
                <div className={StyleHistory.Body}>
                    <Select
                        label="Select record"
                        className="max-w-full"
                        onChange={(e => {
                            let ts = e.target.value
                            if (ts.length === 0) {
                                 return
                            }
                            setSearchParameters(e.target.value)
                        })}
                    >
                        {
                            history.length !== 0 && history.value.map(v => (
                                <SelectItem key={v.name}>
                                    {v.name}
                                </SelectItem>
                            ))
                        }
                    </Select>

                    <Button onClick={remove} isLoading={loader} color="warning" variant="flat" className="w-full">
                        Cleanup history
                    </Button>
                </div>
            </div>
        </Wrapper>
    )
}
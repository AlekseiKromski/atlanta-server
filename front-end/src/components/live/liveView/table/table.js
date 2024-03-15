import {
    Table as TableReact,
    TableBody,
    TableCell,
    TableColumn,
    TableHeader,
    TableRow, useDisclosure
} from "@nextui-org/react";
import TableStyle from "./table.module.css";
import Wrapper from "../../../common/wrapper/wrapper";
export default function Table({isDark, device, datapoints}) {

    return (
        <Wrapper
            title={"Table"}
            modal={{
                title: "Table help",
                body: (
                    <p>
                        TODO: ...
                    </p>
                )
            }}
            isDark={isDark}
        >
            <div className={TableStyle.Body}>
                {
                    datapoints &&  <TableReact className={TableStyle.Table} aria-label="Example static collection table">
                        <TableHeader>
                            <TableColumn>Device ID</TableColumn>
                            <TableColumn>Label</TableColumn>
                            <TableColumn>Value</TableColumn>
                            <TableColumn>Unit</TableColumn>
                            <TableColumn>Flags</TableColumn>
                            <TableColumn>Measurement time</TableColumn>
                            <TableColumn>Created</TableColumn>
                        </TableHeader>
                        <TableBody emptyContent={"No data received"} items={datapoints}>
                            {datapoints.length !== 0 && datapoints.filter(dp => dp.device_id === device).map(dp => (
                                <TableRow key={dp.id}>
                                    <TableCell>{dp.device_id}</TableCell>
                                    <TableCell>{dp.label}</TableCell>
                                    <TableCell>{dp.value}</TableCell>
                                    <TableCell>{dp.unit}</TableCell>
                                    <TableCell>{dp.flags}</TableCell>
                                    <TableCell>{dp.measurement_time}</TableCell>
                                    <TableCell>{dp.created_at}</TableCell>
                                </TableRow>
                            ))}
                        </TableBody>
                    </TableReact>
                }
            </div>
        </Wrapper>
    )
}
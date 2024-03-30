import {Card, CardBody, CardFooter, CardHeader, Chip, Divider, Link} from "@nextui-org/react";
import AlertStyle from "./alert.module.css"

export default function Alert({type, description, link}) {
    return (
        <Card className={AlertStyle.Alert + " " + AlertStyle.Show + " max-w-[400px]"}>
            <CardHeader className="flex gap-3">
                <div className="flex justify-between w-full">
                    {
                        type && type === "warning" && <Chip color="warning">Warning</Chip> ||
                        type && type === "error" && <Chip color="danger">Error</Chip> ||
                        type && type === "info" && <Chip color="default">Info</Chip>
                    }
                </div>
            </CardHeader>
            <Divider/>
            <CardBody>
                <p>{description}</p>
            </CardBody>
            <Divider/>
            <CardFooter>
                <Link
                    isExternal
                    showAnchorIcon
                    href={link}
                >
                    Check official documentation page
                </Link>
            </CardFooter>
        </Card>
    )
}
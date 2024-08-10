import Link from "next/link"
import {redirect} from 'next/navigation'
import {Button} from "@/components/ui/button"
import {Table, TableHeader, TableRow, TableHead, TableBody, TableCell} from "@/components/ui/table"
import {Badge} from "@/components/ui/badge";
import {
    CloudIcon,
    FolderArchiveIcon,
    PlusIcon,
    XIcon,
    LogOutIcon,
    UploadIcon,
    DownloadIcon,
    FileTypeIcon
} from "@/components/icons";
import {decodeString} from "@/components/utils/encoding";
import {getFiles, getPath, resolveStorageID} from "@/components/api/file";

function isDir(file: { type: string; }) {
    return file.type === 'DIR';
}

// @ts-ignore
export async function FileExplorer({storageName, path}) {
    const compareFolderPriority = (a: { name: string, type: string; }, b: { name: string, type: string; }) => {
        if (isDir(a) && !isDir(b)) {
            return -1;
        }

        if (!isDir(a) && isDir(b)) {
            return 1;
        }

        return a.name.localeCompare(b.name);
    }

    const bytesToSize = (bytes: number) => {
        const sizes = ['Bytes', 'KB', 'MB', 'GB', 'TB'];
        if (bytes == 0) return '0 Byte';
        const i = parseInt(String(Math.floor(Math.log(bytes) / Math.log(1024))));
        return Math.round(bytes / Math.pow(1024, i)) + ' ' + sizes[i];
    }

    const decodedStorageName = decodeString(storageName);

    const storageID = await resolveStorageID(decodedStorageName);
    const data = await getFiles(storageID, path ? path : 0);
    const files = data || [];
    const pathData = path ? await getPath(storageID, path) : '/';

    return (
        <>
            <div className="flex items-center justify-between mb-6">
                <div>
                    <h1 className="text-2xl font-bold">Files in {decodedStorageName}</h1>
                    <p className="text-muted-foreground text-sm py-1">{pathData}</p>
                </div>
                <div className="flex items-center gap-4">
                    <Link href={`/s/${storageName}/u`}>
                        <Button variant="outline">
                            <UploadIcon className="mr-2 h-4 w-4"/>
                            Upload
                        </Button>
                    </Link>
                </div>
            </div>
            <div className="overflow-auto">
                <Table>
                    <TableHeader>
                        <TableRow>
                            <TableHead>Name</TableHead>
                            <TableHead>Type</TableHead>
                            <TableHead>Last Modified</TableHead>
                            <TableHead className="text-right">Size</TableHead>
                        </TableRow>
                    </TableHeader>
                    <TableBody>
                        {[...files]
                            .sort((a, b) => compareFolderPriority(a, b))
                            .map((file) => (
                                <TableRow key={file.name}>
                                    <TableCell>
                                        <Link href={file.type !== 'DIR' ? `/api/file/${file.id}` : `/s/${storageName}?path=${file.id}`} className="flex items-center gap-2" prefetch={false}>
                                            {file.type === 'DIR' ? <FolderArchiveIcon className="h-5 w-5 text-muted-foreground"/> : <FileTypeIcon type={file.type} className="h-5 w-5 text-muted-foreground"/>}
                                            <span>{file.name}</span>
                                        </Link>
                                    </TableCell>
                                    <TableCell>
                                        <Badge variant="outline" className="px-2 py-1">
                                            {file.type}
                                        </Badge>
                                    </TableCell>
                                    <TableCell>{file.last_modified.split('T')[0]}</TableCell>
                                    <TableCell className="text-right">{file.type == 'DIR' ? '-' : bytesToSize(file.size)}</TableCell>
                                </TableRow>
                            ))}
                    </TableBody>
                </Table>
            </div>
        </>
    )
}

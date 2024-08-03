import {FileExplorer} from "@/components/file-explorer";
import {Sidebar} from "@/components/sidebar";

function getData() {
    const files = [
        {name: 'file1', type: 'JPEG', modifiedAt: '2021-01-01', size: "21 MB"},
        {name: 'file2', type: 'PPTX', modifiedAt: '2021-01-01', size: "21 MB"},
        {name: 'file3', type: 'MP3', modifiedAt: '2021-01-01', size: "21 MB"},
    ]

    return files;
}

export default function Page({params}: { params: { storage: string } }) {
    return (
        <>
            <Sidebar encodedStorageName={params.storage} />
            <div className="flex flex-1 flex-col px-8 py-10">
                <FileExplorer files={getData()} storageName={params.storage}/>
            </div>
        </>
    );
}

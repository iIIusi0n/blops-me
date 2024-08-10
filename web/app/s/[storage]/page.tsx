import {FileExplorer} from "@/components/file-explorer";
import {Sidebar} from "@/components/sidebar";

export default function Page({params, searchParams}: { params: { storage: string }, searchParams: { [key: string]: string | string[] | undefined } }) {
    return (
        <>
            <Sidebar encodedStorageName={params.storage} />
            <div className="flex flex-1 flex-col px-8 py-10">
                <FileExplorer storageName={params.storage} path={searchParams.path}/>
            </div>
        </>
    );
}

import FileUpload from "@/components/file-upload";
import {Sidebar} from "@/components/sidebar";

export default function Page({params}: { params: { storage: string } }) {
    return (
        <>
            <Sidebar encodedStorageName={params.storage} />
            <div className="flex flex-1 flex-col px-8 py-10">
                <FileUpload storageName={params.storage}/>
            </div>
        </>
    );
}

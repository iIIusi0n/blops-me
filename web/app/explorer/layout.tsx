import {Sidebar} from "@/components/sidebar";

// @ts-ignore
export default function Layout({ children }) {
    return (
        <div className="flex min-h-screen w-full">
            <Sidebar/>
            <div className="flex flex-1 flex-col px-8 py-10">
                {children}
            </div>
        </div>
    )
}
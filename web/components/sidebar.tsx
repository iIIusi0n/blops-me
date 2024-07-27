import Link from "next/link";
import {Button} from "@/components/ui/button";
import {Input} from "@/components/ui/input";
import {CloudIcon, FolderArchiveIcon, XIcon, PlusIcon, LogOutIcon} from "@/components/icons";

function getData() {
    const data = [
        {id: 1, name: "Documents"},
        {id: 2, name: "Images"},
        {id: 3, name: "Videos"},
        {id: 4, name: "Music"},
        {id: 5, name: "Archives"},
    ];

    return data;
}

export function Sidebar() {
    const categories = getData();

    return (
        <div className="hidden w-64 flex-col border-r bg-card p-6 md:flex">
            <Link href="/storage" className="flex items-center gap-2 mb-6" prefetch={false}>
                <CloudIcon className="h-6 w-6"/>
                <span className="text-lg font-semibold">Blops.me</span>
            </Link>
            <nav className="flex-1 space-y-2">
                {categories.map((category) => (
                    <div key={category.id} className="flex items-center justify-between">
                        <Link href="#" className="flex items-center gap-2 rounded-md px-3 py-2 hover:bg-muted"
                              prefetch={false}>
                            <FolderArchiveIcon className="h-5 w-5 text-muted-foreground"/>
                            <span>{category.name}</span>
                        </Link>
                        <Button variant="ghost" size="icon">
                            <XIcon className="h-4 w-4 text-muted-foreground"/>
                            <span className="sr-only">Remove {category.name}</span>
                        </Button>
                    </div>
                ))}
                <div className="flex items-center gap-2 py-5">
                    <Input
                        type="text"
                        placeholder="Add new category"
                        name="category"
                        className="rounded-md bg-background px-3 py-2 text-sm shadow-sm transition-colors focus:outline-none focus:ring-1 focus:ring-ring disabled:cursor-not-allowed disabled:opacity-50"
                    />
                    <Button variant="ghost" size="icon">
                        <PlusIcon className="h-4 w-4 text-muted-foreground"/>
                        <span className="sr-only">Add new category</span>
                    </Button>
                </div>
            </nav>
            <div className="mt-auto flex items-center gap-2">
                <Button variant="ghost" size="icon">
                    <LogOutIcon className="h-4 w-4 text-muted-foreground"/>
                    <span className="sr-only">Logout</span>
                </Button>
            </div>
        </div>
    );
}

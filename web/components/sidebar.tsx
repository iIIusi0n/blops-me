import Link from "next/link";
import {Button} from "@/components/ui/button";
import {Input} from "@/components/ui/input";
import {
    CloudIcon,
    FolderArchiveIcon,
    LogOutIcon,
    PlusIcon,
    XIcon
} from "@/components/icons";
import {decodeString, encodeString} from "@/components/utils/encoding";
import {createStorage, deleteStorage, getStorages, Storage} from "@/components/api/storage";
import {redirect} from "next/navigation";
import {revalidatePath} from "next/cache";
import {cookies} from "next/headers";

export async function Sidebar({ encodedStorageName } : { encodedStorageName: string }) {
    const {storages} = await getStorages();

    if (storages.length > 0 && encodedStorageName === "") {
        redirect(`/s/${encodeString(storages[0].name)}`);
    }

    try {
        const storageName = decodeString(encodedStorageName);
        if (storages.findIndex(storage => storage.name === storageName) === -1) {
            throw new Error("Storage not found");
        }
    } catch (e) {
        if (encodedStorageName !== "") {
            redirect("/s");
        }
    }

    const handleAddStorage = async (formData: FormData) => {
        "use server";
        const newStorageName = formData.get('storage');

        if (!newStorageName || newStorageName === '') {
            return;
        }

        await createStorage(newStorageName as string);
        revalidatePath(`/s/${encodeString((newStorageName as string).toUpperCase())}`);

        formData.set('storage', '');

        redirect(`/s/${encodeString((newStorageName as string).toUpperCase())}`);
    }

    const handleDeleteStorage = async (deletedStorageID: number, deletedStorageName: string) => {
        "use server";
        await deleteStorage(deletedStorageID);

        const storageName = decodeString(encodedStorageName);
        let redirectPath = "/s";
        if (storageName !== deletedStorageName) {
            redirectPath += `/${encodedStorageName}`;
        }

        revalidatePath(redirectPath);
        redirect(redirectPath);
    }

    const handleLogout = async () => {
        "use server";
        cookies().delete('token');
        redirect("/");
    }

    return (
        <>
            <div className="hidden w-64 flex-col border-r bg-card p-6 md:flex">
                <Link href="/s" className="flex items-center gap-2 mb-6"
                      prefetch={false}>
                    <CloudIcon className="h-6 w-6"/>
                    <span className="text-lg font-semibold">Blops.me</span>
                </Link>
                <nav className="flex-1 space-y-2">
                    {storages.map((storage) => (
                        <div key={storage.id}
                             className="flex items-center justify-between">
                            <Link href={`/s/${encodeString(storage.name)}`}
                                  className="flex items-center gap-2 rounded-md px-3 py-2 hover:bg-muted"
                                  prefetch={false}>
                                <FolderArchiveIcon
                                    className="h-5 w-5 text-muted-foreground"/>
                                <span>{storage.name}</span>
                            </Link>
                            <form action={handleDeleteStorage.bind(null, storage.id, storage.name)}>
                                <Button variant="ghost" size="icon">
                                    <XIcon
                                        className="h-4 w-4 text-muted-foreground"/>
                                    <span
                                        className="sr-only">Remove {storage.name}</span>
                                </Button>
                            </form>
                        </div>
                    ))}
                    <div className="flex items-center gap-2 py-2">
                        <form action={handleAddStorage}>
                            <div className="flex items-center gap-2 py-2">
                                <Input
                                    type="text"
                                    placeholder="Add new storage"
                                    name="storage"
                                    maxLength={16}
                                    className="rounded-md bg-background px-3 py-2 text-sm shadow-sm transition-colors focus:outline-none focus:ring-1 focus:ring-ring disabled:cursor-not-allowed disabled:opacity-50"
                                />
                                <Button variant="ghost" size="icon"
                                        type="submit">
                                    <PlusIcon
                                        className="h-4 w-4 text-muted-foreground"/>
                                    <span
                                        className="sr-only">Add new storage</span>
                                </Button>
                            </div>
                        </form>
                    </div>
                </nav>
                <div className="mt-auto flex items-center gap-2">
                    <form action={handleLogout}>
                        <Button variant="ghost" size="icon">
                            <LogOutIcon className="h-4 w-4 text-muted-foreground"/>
                            <span className="sr-only">Logout</span>
                        </Button>
                    </form>
                </div>
            </div>
        </>
    );
}

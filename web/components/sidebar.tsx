"use client";

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
import {useEffect, useState} from "react";
import {encodeString} from "@/components/utils/encoding";
import {createStorage, deleteStorage, getStorages, Storage} from "@/components/api/storage";
import {usePathname, useRouter} from "next/navigation";


export function Sidebar() {
    const path = usePathname();
    const router = useRouter();

    const [storages, setStorages] = useState<Storage[]>([]);
    const [newStorageName, setNewStorageName] = useState("");

    useEffect(() => {
        getStorages().then((data) => setStorages(data.storages));
    }, []);

    useEffect(() => {
        if (storages.length > 0 && path === '/s') {
            router.push(`/s/${encodeString(storages[0].name)}`);
        }
    }, [storages, path]);

    const handleAddStorage = async () => {
        if (newStorageName) {
            await createStorage(newStorageName);
            const updatedData = await getStorages();
            setStorages(updatedData.storages);
            setNewStorageName('');
            router.push(`/s/${encodeString(newStorageName.toUpperCase())}`);
        }
    };

    const logout = async () => {
        await fetch('/auth/logout');

        router.push('/');
    }

    const handleDeleteStorage = async (id: number) => {
        const originalStorageName = storages.find(storage => storage.id === id)?.name || '';

        await deleteStorage(id);
        const updatedData = await getStorages();
        setStorages(updatedData.storages);

        if (updatedData.storages.length === 0) {
            router.push('/s');
            return;
        }

        if (path === `/s/${encodeString(originalStorageName)}`) {
            router.push(`/s/${encodeString(updatedData.storages[0].name)}`);
        }
    };

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
                            <Button variant="ghost" size="icon"
                                    onClick={() => handleDeleteStorage(storage.id)}>
                                <XIcon
                                    className="h-4 w-4 text-muted-foreground"/>
                                <span
                                    className="sr-only">Remove {storage.name}</span>
                            </Button>
                        </div>
                    ))}
                    <div className="flex items-center gap-2 py-2">
                        <Input
                            type="text"
                            placeholder="Add new storage"
                            name="storage"
                            value={newStorageName}
                            maxLength={16}
                            onChange={(e) => setNewStorageName(e.target.value)}
                            className="rounded-md bg-background px-3 py-2 text-sm shadow-sm transition-colors focus:outline-none focus:ring-1 focus:ring-ring disabled:cursor-not-allowed disabled:opacity-50"
                        />
                        <Button variant="ghost" size="icon"
                                onClick={handleAddStorage}>
                            <PlusIcon
                                className="h-4 w-4 text-muted-foreground"/>
                            <span className="sr-only">Add new storage</span>
                        </Button>
                    </div>
                </nav>
                <div className="mt-auto flex items-center gap-2">
                    <Button variant="ghost" size="icon" onClick={logout}>
                        <LogOutIcon className="h-4 w-4 text-muted-foreground"/>
                        <span className="sr-only">Logout</span>
                    </Button>
                </div>
            </div>
        </>
    );
}

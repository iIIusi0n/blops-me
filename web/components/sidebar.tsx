"use client";

import Link from "next/link";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { CloudIcon, FolderArchiveIcon, XIcon, PlusIcon, LogOutIcon } from "@/components/icons";
import { useState, useEffect } from "react";

async function createStorage(storageName: string) {
    await fetch("http://localhost:8080/api/storage", {
        method: "POST",
        credentials: 'include',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({'storage_name': storageName})
    });
}

async function deleteStorage(storageID: number) {
    await fetch(`http://localhost:8080/api/storage`, {
        method: "DELETE",
        credentials: 'include',
        headers: {
            'storage-id': storageID.toString(),
        }
    });
}

async function getStorages(): Promise<any> {
    const resp = await fetch("http://localhost:8080/api/storage", {
        credentials: 'include',
    });
    if (!resp.ok) {
        return [];
    }

    return resp.json();
}

export function Sidebar() {
    const [storages, setStorages] = useState<any[]>([]);
    const [newStorageName, setNewStorageName] = useState("");

    useEffect(() => {
        getStorages().then((data) => setStorages(data.storages));
    }, []);

    const handleAddStorage = async () => {
        if (newStorageName) {
            await createStorage(newStorageName);
            const updatedData = await getStorages();
            setStorages(updatedData.storages);
            setNewStorageName('');
        }
    };

    const handleDeleteStorage = async (id: number) => {
        await deleteStorage(id);
        const updatedData = await getStorages();
        setStorages(updatedData.storages);
    };

    return (
        <div className="hidden w-64 flex-col border-r bg-card p-6 md:flex">
            <Link href="/s" className="flex items-center gap-2 mb-6" prefetch={false}>
                <CloudIcon className="h-6 w-6"/>
                <span className="text-lg font-semibold">Blops.me</span>
            </Link>
            <nav className="flex-1 space-y-2">
                {storages.map((storage) => (
                    <div key={storage.id} className="flex items-center justify-between">
                        <Link href="#" className="flex items-center gap-2 rounded-md px-3 py-2 hover:bg-muted"
                              prefetch={false}>
                            <FolderArchiveIcon className="h-5 w-5 text-muted-foreground"/>
                            <span>{storage.name}</span>
                        </Link>
                        <Button variant="ghost" size="icon" onClick={() => handleDeleteStorage(storage.id)}>
                            <XIcon className="h-4 w-4 text-muted-foreground"/>
                            <span className="sr-only">Remove {storage.name}</span>
                        </Button>
                    </div>
                ))}
                <div className="flex items-center gap-2 py-5">
                    <Input
                        type="text"
                        placeholder="Add new storage"
                        name="storage"
                        value={newStorageName}
                        maxLength={16}
                        onChange={(e) => setNewStorageName(e.target.value)}
                        className="rounded-md bg-background px-3 py-2 text-sm shadow-sm transition-colors focus:outline-none focus:ring-1 focus:ring-ring disabled:cursor-not-allowed disabled:opacity-50"
                    />
                    <Button variant="ghost" size="icon" onClick={handleAddStorage}>
                        <PlusIcon className="h-4 w-4 text-muted-foreground"/>
                        <span className="sr-only">Add new storage</span>
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

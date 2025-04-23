"use client";

import { useSession, signIn, signOut } from "next-auth/react";
import { useState } from "react";

// Define a type for the group object
interface Group {
    id: string;
    // Add other properties as needed
}

export default function Home() {
    const { data: session } = useSession();
    const [groups, setGroups] = useState<Group[]>([]);

    async function fetchGroup() {
        console.log(session)
        try {
            // Use credentials: 'include' to ensure cookies are sent
            const response = await fetch("http://localhost:8000/api/v1/applications-groups/", {
                headers: {
                    "Authorization": `Bearer ${session?.accessToken || ""}`,
                },
                credentials: 'include' // This ensures cookies are sent with the request
            });

            if (!response.ok) throw new Error('Network response was not ok');
            const fetchedGroups = await response.json();
            setGroups(fetchedGroups.data || []);
        } catch (error) {
            console.error("Failed to fetch groups:", error);
        }
    }

    return (
        <div className="p-6 text-center">
            {session ? (
                <>
                    <h1 className="text-xl mb-4">Hello, {session.user?.name}</h1>
                    <p>{session.user?.email}</p>

                    <button
                        onClick={() => fetchGroup()}
                        className="bg-red-500 text-white px-4 py-2 rounded"
                    >
                        Groups
                    </button>

                    <ul>
                        {
                            groups.length > 0 && groups.map((group: Group) => (
                                <li key={group.id}>{group.id}</li>
                            ))
                        }
                    </ul>

                    <button
                        onClick={() => signOut()}
                        className="bg-red-500 text-white px-4 py-2 rounded"
                    >
                        Sign Out
                    </button>
                </>
            ) : (
                <button
                    onClick={() => signIn("keycloak")}
                    className="bg-blue-500 text-white px-4 py-2 rounded"
                >
                    Sign In with Keycloak
                </button>
            )}
        </div>
    );
}
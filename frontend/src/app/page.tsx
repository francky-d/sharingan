"use client";

import { useSession, signIn, signOut } from "next-auth/react";
import { useEffect, useState } from "react";

export default function Home() {
  const { data: session } = useSession();
  const [groups, setGroups] = useState([]);

   async function fetchGroup(){
    const fetchedGroups = await fetch("http://localhost:8000/api/v1/applications-groups")
     .then(res => res.json())
        .catch(error => console.log(error));
     setGroups(fetchedGroups);
  }
  return (
    <div className="p-6 text-center">
      {session ? (
        <>
          <h1 className="text-xl mb-4">Hello, {session.user?.name}</h1>
          <p>{session.user?.id}</p>
          <p>{session.user?.email}</p>

          <button
              onClick={() => fetchGroup()}
              className="bg-red-500 text-white px-4 py-2 rounded"
          >
           Groups
          </button>

          <ul>
            {

                   groups?.length > 0 && groups.map((group: any) => (
                   <li key={group.id}>{group?.id}</li>
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
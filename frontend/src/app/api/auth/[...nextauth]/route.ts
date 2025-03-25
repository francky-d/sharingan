import NextAuth from "next-auth";
import KeycloakProvider from "next-auth/providers/keycloak";

// Create the handler with your configuration
const handler = NextAuth({
    providers: [
        KeycloakProvider({
            clientId: process.env.KEYCLOAK_CLIENT_ID,
            clientSecret: process.env.KEYCLOAK_CLIENT_SECRET || "",
            issuer: `${process.env.KEYCLOAK_URL}/realms/${process.env.KEYCLOAK_REALM}`,
        }),
    ],
    callbacks: {
        async jwt({ token, account }) {
            if (account) token.accessToken = account.access_token;
            return token;
        },
        async session({ session, token }) {
            session.accessToken = token.accessToken;
            return session;
        },
    },
    secret: process.env.NEXTAUTH_SECRET,
});

// Export the handler as named exports for GET and POST
export const GET = handler;
export const POST = handler;

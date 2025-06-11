import { ApplicationGroups } from "./application-groups"
import { Incident } from "./Incident"

export interface Application {
    UserID: Number,
    ApplicationGroupID: Number,
    Url: string
    UrlToWatch: string
    HttpSuccessCode: Number,
    ApplicationGroup: ApplicationGroups
    Incidents: Incident[]

}
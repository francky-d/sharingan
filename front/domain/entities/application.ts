import { ApplicationGroup } from "./application-group"
import { Incident } from "./Incident"

export interface Application {
    UserID: Number,
    ApplicationGroupID: Number,
    Url: string
    UrlToWatch: string
    HttpSuccessCode: Number,
    ApplicationGroup: ApplicationGroup
    Incidents: Incident[]

}
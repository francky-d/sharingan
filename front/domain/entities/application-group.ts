import { Application } from "./application"

export interface ApplicationGroup {
    ID: Number,
    Name: string,
    Applications: Application[]
}
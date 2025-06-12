import { Application } from "./application"

export interface ApplicationGroup {
    ID: number,
    Name: string,
    Applications: Application[]
}
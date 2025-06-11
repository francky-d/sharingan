import { Incident } from "../entities/Incident";

export interface IncidentRespositoryInterface {
    getIncidents(): Promise<Incident[]>;
    getIncidentsByID(ID: Number): Promise<Incident>;
    getIncidentByApplicationID(appID: Number): Promise<Incident[]>;
    createIncident(incident: Incident, appID: Number): Promise<Incident>
    updateIncident(incident: Incident): Promise<Incident>;
    deleteIncident(ID: Number): Promise<void>
}
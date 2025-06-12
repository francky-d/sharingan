import ApplicationGroupRepositoryInterface from "../../domain/ports/ApplicationGroupRepository";
import {ApplicationGroup} from "../../domain/entities/application-group";
import axios from "axios";

export default class ApplicationGroupApi implements ApplicationGroupRepositoryInterface {
    apiUrl: string = "/api/v1/application-groups";

    async getApplicationGroups(): Promise<ApplicationGroup[]> {
        const response = await axios.get<ApplicationGroup[]>(this.apiUrl);
        return response.data
    }

    async getApplicationGroupByID(ID: Number): Promise<ApplicationGroup> {
        const response = await axios.get<ApplicationGroup>(`${this.apiUrl}/${ID}`);
        return response.data;
    }

    async createApplicationGroup(appGroup: ApplicationGroup): Promise<ApplicationGroup> {
        const response = await axios.post<ApplicationGroup>(this.apiUrl, appGroup);
        return response.data;
    }

    async updateApplicationGroup(appGroup: ApplicationGroup): Promise<ApplicationGroup> {
        const response = await axios.put<ApplicationGroup>(`${this.apiUrl}/${appGroup.id}`, appGroup);
        return response.data;
    }

    async deleteApplicationGroup(ID: Number): Promise<void> {
        await axios.delete(`${this.apiUrl}/${ID}`);
    }

}
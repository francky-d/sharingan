import { configureStore } from '@reduxjs/toolkit'
import {applicationReducer} from "../feature/application/applicationSlice";

 function globalReducer() {
    return  {
         application : applicationReducer,
     }
 }






export const store = configureStore({
    reducer: globalReducer()
})

export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch
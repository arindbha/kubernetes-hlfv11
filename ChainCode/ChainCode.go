package main

import (
    "fmt"
    "github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
    "encoding/json"
)

type Patient struct {
    id string `json:"id"`
    name string `json:"name"`
    age string `json:"age"`
    doctorId string `json:"String"`
    doctorName string `json:"String"`
}

func (t *Patient) Init(stub shim.ChaincodeStubInterface) pb.Response {
    _,args:= stub.GetFunctionAndParameters()
    if len(args)<0 {
        return shim.Error("Invalid arguements! Expecting value")
    }
    return shim.Success([]byte("Chaincode Initialized"))
}

func (t *Patient) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	//myLogger.Debug("Invoke Chaincode...")
	function,_:= stub.GetFunctionAndParameters()
	if function == "registerPatient" {
		// Insert Patient details
		return t.registerPatient(stub)
	} else if function == "getPatient" {
		// Allocates slot for the asset
		return t.getPatient(stub)
	}
	return shim.Error("Invalid invoke function name. Expecting \"invoke\" \"delete\" \"query\"")
}

func (t *Patient) registerPatient(stub shim.ChaincodeStubInterface) pb.Response {
    _,args:=stub.GetFunctionAndParameters()
    if len(args)<5 {
        return shim.Error("Invalid Arguement!")
    }
    //var patient_age int
    patient_id := args[0]
    patient_name := args[1]
     patient_age:= args[2]
    doctor_id := args[3]
    doctor_name := args[4]
    PatientasByte,err := stub.GetState(patient_id)
    if err!=nil {
        return shim.Error("Failed to get Patient Details")
    } else if PatientasByte !=nil {
        return shim.Error ("Patient already Exists : "+patient_id)
    }
    //doctor := &Doctor{doctor_id,doctor_name}
    patient:= &Patient{patient_id,patient_name,patient_age,doctor_id,doctor_name}
    patient_new,_:= json.Marshal(patient)
    err=stub.PutState(patient_id,patient_new)
    if err!=nil {
        return shim.Error("error_62")
    }
    return shim.Success([]byte("Successfully registration Completed!"))
}

func (t *Patient) getPatient(stub shim.ChaincodeStubInterface) pb.Response {
    //var err error
    _,args:= stub.GetFunctionAndParameters()
    if len(args)<1 {
        return shim.Error("Invalid Arguements")
    }
    patientId:=args[0]
    patient_details,_:= stub.GetState(patientId)
    if patient_details == nil {
        return shim.Error("Patient doesn't exist")
    }
    patientJson,_:=json.Marshal(patient_details)
    return shim.Success([]byte(patientJson))
}

func main() {
    err:= shim.Start(new(Patient))
    if err != nil {
		fmt.Printf("Error starting Medical: %s", err)
	}
}


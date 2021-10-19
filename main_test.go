package main

import (
	"bytes"
	"os"
	"strings"
	"super-code-gen/pkg/generate"
	"testing"
)

func TestGenRoute(t *testing.T) {
	getUserFileWant :=
		[]byte(`import { RequestHandler } from "express";

interface GetUserReqBody {

}

export const getUser:RequestHandler=async (req,res)=>{
const body = req.body as GetUserReqBody
}`)
	addUserFileWant :=
		[]byte(`import { RequestHandler } from "express";

interface AddUserReqBody {

}

export const addUser:RequestHandler=async (req,res)=>{
const body = req.body as AddUserReqBody
}`)
	indexFileWant := []byte(`export { getUser } from "./getUser";
export { addUser } from "./addUser";
`)
	userRoutesFileWant := []byte(`import { Router } from "express";
import { getUser, addUser, } from "../controllers/user-controllers";

import { body } from "express-validator";

const userRouter = Router();
userRouter.post("/getUser", getUser);
userRouter.post("/addUser", addUser);

export default user;`)
	name := "user"
	routes := []string{"user/getUser", "user/addUser"}
	validate, body := true, true
	paths := make([]string, 0)
	for _, arg := range routes {
		paths = append(paths, strings.Split(arg, "/")[1])
	}
	generate.GenRoute(name, paths, validate)
	generate.GenController(name, paths, body)
	controllerRoot := "controllers/user-controllers/"
	getUserFile, e := os.ReadFile(controllerRoot + "getUser.ts")
	addUserFile, _ := os.ReadFile(controllerRoot + "addUser.ts")
	indexFile, _ := os.ReadFile(controllerRoot + "index.ts")
	userRoutesFile, _ := os.ReadFile("routes/user-routes.ts")
	files := [][][]byte{
		{getUserFile, getUserFileWant},
		{addUserFile, addUserFileWant},
		{indexFile, indexFileWant},
		{userRoutesFile, userRoutesFileWant},
	}

	for _, fileGroup := range files {
		res := bytes.Compare(fileGroup[0], []byte(fileGroup[1]))
		if res == 1 {
			t.Fail()
		}
		if e != nil {
			t.Fatalf("File: getUser, Error %v", e)
		}
	}

	os.RemoveAll(controllerRoot)
	os.Remove("routes/user-routes.ts")

}

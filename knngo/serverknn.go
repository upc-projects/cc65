package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"
	"math/rand"
	"net"
	"os"
	"sort"
	"strconv"
	"strings"
)

var addrs []string
var myhost string

// var k int = 3

func main() {

	gin := bufio.NewReader(os.Stdin)
	fmt.Print("Ingrese host: ")
	myhost, _ = gin.ReadString('\n')
	myhost = strings.TrimSpace(myhost)
	myip := fmt.Sprintf("localhost:%s", myhost)

	fmt.Printf("Soy %s\n", myip)
	go registerServer(myhost)
	go hotServer(myhost)

	gin = bufio.NewReader(os.Stdin)
	fmt.Print("Ingrese direccion remota: ")
	remoteHost, _ := gin.ReadString('\n')
	remoteHost = strings.TrimSpace(remoteHost)
	remoteIP := fmt.Sprintf("localhost:%s", remoteHost)

	if remoteIP != "" {
		registerSend(remoteHost, myhost)
	}

	go func() {
		fmt.Print("Ingrese el valor de k: ")
		strNum, _ := gin.ReadString('\n')
		if strNum != "" {
			num, _ := strconv.Atoi(strings.TrimSpace(strNum))
			hotSend(num)
		}
	}()

	notifyServer(myhost)
}

func hotServer(hostAddr string) {
	newhost, _ := strconv.Atoi(strings.TrimSpace(hostAddr))
	newhost = newhost + 3
	host := fmt.Sprintf("localhost:%d", newhost)
	ln, _ := net.Listen("tcp", host)
	defer ln.Close()
	for {
		conn, _ := ln.Accept()
		go handleHot(conn)
	}
}

func handleHot(conn net.Conn) {
	defer conn.Close()

	r := bufio.NewReader(conn)
	strNum, _ := r.ReadString('\n')
	num, _ := strconv.Atoi(strings.TrimSpace(strNum))
	fmt.Printf("Recibimos el %d\n", num)
	if num == 0 {
		conn.Close()
	} else {
		hotSend(num - 1)
	}
}

func hotSend(num int) {
	if num == 0 {
		num = 3
	}
	idx := rand.Intn(len(addrs))
	fmt.Println(idx)
	// fmt.Printf("Enviando %d a %s\n", num, addrs[idx])
	irisData, err := os.Open("iris.csv")
	errHandle(err)
	defer irisData.Close()

	reader := csv.NewReader(irisData)
	reader.Comma = ','

	var recordSet []irisRecord

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		errHandle(err)

		recordSet = append(recordSet, parseIrisRecord(record))
	}

	var testSet []irisRecord
	var trainSet []irisRecord

	for i := range recordSet {
		if rand.Float64() < 0.05 {
			trainSet = append(trainSet, recordSet[i])
		} else {
			testSet = append(testSet, recordSet[i])
		}
	}

	var predictions []string

	for x := range testSet {
		neighbors := getNeighbors(trainSet, testSet[x], num)
		fmt.Println(neighbors)
		result := getResponse(neighbors)
		predictions = append(predictions, result[0].key)
		// fmt.Printf("Predicted: %s, Actual: %s\n", result[0].key, testSet[x].species)
	}

	accuracy := getAccuracy(testSet, predictions)
	fmt.Printf("Accuracy: %f%s\n", accuracy, "%")

	host, _ := strconv.Atoi(strings.TrimSpace(addrs[idx]))
	host = host + 3
	remote := fmt.Sprintf("localhost:%d", host)
	conn, _ := net.Dial("tcp", remote)
	defer conn.Close()
	fmt.Fprintln(conn, num)
}

func registerSend(remoteAddr, hostAddr string) {

	remotehost, _ := strconv.Atoi(strings.TrimSpace(remoteAddr))
	remotehost = remotehost + 1
	remote := fmt.Sprintf("localhost:%d", remotehost)
	conn, _ := net.Dial("tcp", remote)
	defer conn.Close()

	// Enviar direccion
	hostAddr = fmt.Sprintf("localhost:%s", hostAddr)
	fmt.Fprintln(conn, hostAddr)

	// Recibir lista de direcciones
	r := bufio.NewReader(conn)
	strAddrs, _ := r.ReadString('\n')
	var respAddrs []string
	json.Unmarshal([]byte(strAddrs), &respAddrs)

	// agregamos direcciones de nodos a propia libreta
	var aux []string
	for _, addr := range respAddrs {
		if addr == strings.TrimSpace(remoteAddr) {
			return
		} else if addr != myhost {
			aux = append(aux, addr)
		}
	}
	addrs = append(aux, strings.TrimSpace(remoteAddr))
	fmt.Println(addrs)
}

func registerServer(hostAddr string) {
	newHost, _ := strconv.Atoi(strings.TrimSpace(hostAddr))
	newHost = newHost + 1

	host := fmt.Sprintf("localhost:%d", newHost)
	ln, _ := net.Listen("tcp", host)
	defer ln.Close()
	for {
		conn, _ := ln.Accept()
		go handleRegister(conn)
	}
}

func handleRegister(conn net.Conn) {
	defer conn.Close()

	// Recibimos addr del nuevo nodo
	r := bufio.NewReader(conn)
	remoteIP, _ := r.ReadString('\n')
	remoteIP = strings.TrimSpace(remoteIP)
	remoteIP = strings.TrimPrefix(remoteIP, "localhost:")

	// respondemos enviando lista de direcciones de nodos actuales
	byteAddrs, _ := json.Marshal(addrs)
	fmt.Fprintf(conn, "%s\n", string(byteAddrs))

	// notificar a nodos actuales de llegada de nuevo nodo
	for _, addr := range addrs {
		notifySend(addr, remoteIP)
	}

	// Agregamos nuevo nodo a la lista de direcciones
	for _, addr := range addrs {
		if addr == strings.TrimSpace(remoteIP) {
			return
		}
	}
	addrs = append(addrs, strings.TrimSpace(remoteIP))
	fmt.Println(addrs)
}

func notifySend(addr, remoteIP string) {
	host, _ := strconv.Atoi(strings.TrimSpace(addr))
	host = host + 2
	remote := fmt.Sprintf("localhost:%d", host)
	conn, _ := net.Dial("tcp", remote)
	defer conn.Close()
	fmt.Fprintln(conn, remoteIP)
}

func notifyServer(hostAddr string) {
	newHost, _ := strconv.Atoi(strings.TrimSpace(hostAddr))
	newHost = newHost + 2
	host := fmt.Sprintf("localhost:%d", newHost)
	ln, _ := net.Listen("tcp", host)
	defer ln.Close()
	for {
		conn, _ := ln.Accept()
		go handleNotify(conn)
	}
}
func handleNotify(conn net.Conn) {
	defer conn.Close()

	// Recibimos addr del nuevo nodo
	r := bufio.NewReader(conn)
	remoteIP, _ := r.ReadString('\n')
	remoteIP = strings.TrimSpace(remoteIP)
	remoteIP = strings.TrimPrefix(remoteIP, "localhost:")

	// Agregamos nuevo nodo a la lista de direcciones
	for _, addr := range addrs {
		if addr == strings.TrimSpace(remoteIP) || strings.TrimSpace(remoteIP) == myhost {
			return
		}
	}
	addrs = append(addrs, strings.TrimSpace(remoteIP))
	fmt.Println(addrs)
}

func myIP() string { // mandrakeando ando
	ifaces, _ := net.Interfaces()
	for _, iface := range ifaces {
		if strings.HasPrefix(iface.Name, "Local") {
			addrs, _ := iface.Addrs()
			for _, addr := range addrs {
				switch v := addr.(type) {
				case *net.IPNet:
					return v.IP.String()
				case *net.IPAddr:
					return v.IP.String()
				}
			}
		}
	}
	return ""
}

// KNN algorithm

type irisRecord struct {
	sepalLength float64
	sepalWidth  float64
	petalLength float64
	petalWidth  float64
	species     string
}

func getAccuracy(testSet []irisRecord, predictions []string) float64 {
	correct := 0

	for x := range testSet {
		if testSet[x].species == predictions[x] {
			correct++
		}
	}

	return (float64(correct) / float64(len(testSet))) * 100.00
}

type classVote struct {
	key   string
	value int
}

type sortedClassVotes []classVote

func (scv sortedClassVotes) Len() int           { return len(scv) }
func (scv sortedClassVotes) Less(i, j int) bool { return scv[i].value < scv[j].value }
func (scv sortedClassVotes) Swap(i, j int)      { scv[i], scv[j] = scv[j], scv[i] }

func getResponse(neighbors []irisRecord) sortedClassVotes {
	classVotes := make(map[string]int)

	for x := range neighbors {
		response := neighbors[x].species
		if contains(classVotes, response) {
			classVotes[response]++
		} else {
			classVotes[response] = 1
		}
	}

	scv := make(sortedClassVotes, len(classVotes))
	i := 0
	for k, v := range classVotes {
		scv[i] = classVote{k, v}
		i++
	}

	sort.Sort(sort.Reverse(scv))
	return scv
}

type distancePair struct {
	record   irisRecord
	distance float64
}

type distancePairs []distancePair

func (slice distancePairs) Len() int           { return len(slice) }
func (slice distancePairs) Less(i, j int) bool { return slice[i].distance < slice[j].distance }
func (slice distancePairs) Swap(i, j int)      { slice[i], slice[j] = slice[j], slice[i] }

func getNeighbors(trainingSet []irisRecord, testRecord irisRecord, k int) []irisRecord {
	var distances distancePairs
	for i := range trainingSet {
		dist := euclidianDistance(testRecord, trainingSet[i])
		distances = append(distances, distancePair{trainingSet[i], dist})
	}

	sort.Sort(distances)

	var neighbors []irisRecord

	for x := 0; x < k; x++ {
		neighbors = append(neighbors, distances[x].record)
	}

	return neighbors
}

func euclidianDistance(instanceOne irisRecord, instanceTwo irisRecord) float64 {
	var distance float64

	distance += math.Pow((instanceOne.petalLength - instanceTwo.petalLength), 2)
	distance += math.Pow((instanceOne.petalWidth - instanceTwo.petalWidth), 2)
	distance += math.Pow((instanceOne.sepalLength - instanceTwo.sepalLength), 2)
	distance += math.Pow((instanceOne.sepalWidth - instanceTwo.sepalWidth), 2)

	return math.Sqrt(distance)
}

func parseIrisRecord(record []string) irisRecord {
	var iris irisRecord

	iris.sepalLength, _ = strconv.ParseFloat(record[0], 64)
	iris.sepalWidth, _ = strconv.ParseFloat(record[1], 64)
	iris.petalLength, _ = strconv.ParseFloat(record[2], 64)
	iris.petalWidth, _ = strconv.ParseFloat(record[3], 64)
	iris.species = record[4]

	return iris
}

func errHandle(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func contains(votesMap map[string]int, name string) bool {
	for s := range votesMap {
		if s == name {
			return true
		}
	}

	return false
}

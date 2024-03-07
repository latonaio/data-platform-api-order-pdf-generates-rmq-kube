package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-api-order-pdf-generates-rmq-kube/DPFM_API_Input_Formatter"
	dpfm_api_output_formatter "data-platform-api-order-pdf-generates-rmq-kube/DPFM_API_Output_Formatter"
	"data-platform-api-order-pdf-generates-rmq-kube/config"
	"encoding/json"
	"fmt"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	lnpdf "github.com/latonaio/golang-pdf-library"
	"io/ioutil"
	"math/rand"
	"os"
	"sync"
	"time"
)

func (c *DPFMAPICaller) process(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	accepter []string,
	errs *[]error,
	log *logger.Logger,
	conf *config.Conf,
) (interface{}, *string) {
	var header *[]dpfm_api_output_formatter.Header
	var mountPath *string

	for _, fn := range accepter {
		switch fn {
		case "Order":
			func() {
				header, mountPath, _ = c.GeneratePDF(input, errs, log, conf)
			}()
		}
	}

	data := &dpfm_api_output_formatter.Message{
		Header: *header,
	}

	return data, mountPath
}

func (c *DPFMAPICaller) GeneratePDF(
	input *dpfm_api_input_reader.SDC,
	errs *[]error,
	log *logger.Logger,
	conf *config.Conf,
) (*[]dpfm_api_output_formatter.Header, *string, error) {
	var data []dpfm_api_output_formatter.Header

	randomString := generateRandomString(10)
	mountPath := ""

	for _, headerData := range input.Header {
		var err error

		err = mkdirP(fmt.Sprintf(
			"%s/%s",
			conf.MountPath,
			randomString,
		))
		if err != nil {
			return nil, nil, err
		}

		outputPath := fmt.Sprintf(
			"%s/%s/%s%s",
			conf.MountPath,
			randomString,
			fmt.Sprintf("output"),
			".pdf",
		)

		mountPath = outputPath

		dataJsonFilePath := fmt.Sprintf("%s/%s", conf.MountPath, "data.json")
		dataJsonFile, err := os.Create(
			dataJsonFilePath,
		)
		if err != nil {
			return nil, nil, err
		}

		err = copyPDF(
			fmt.Sprintf("%s/%s", conf.MaterialPath, "layout.pdf"),
			outputPath,
		)
		if err != nil {
			return nil, nil, err
		}

		// TODO 後でDBから取得することになりそう
		templateFileName := fmt.Sprintf("%s/%s", conf.MaterialPath, "template.json")

		pdfData := *dpfm_api_output_formatter.SetToPdfData(
			&headerData,
			input.Items,
			input.ItemPricingElements,
			input.Partners,
		)

		pdfDataJsonBytes, err := json.Marshal(pdfData)
		if err != nil {
			return nil, nil, err
		}

		_, err = dataJsonFile.WriteString(string(pdfDataJsonBytes))
		if err != nil {
			return nil, nil, err
		}

		lnpdf.Builder{
			TemplatePath:   templateFileName,
			DataSourcePath: dataJsonFilePath,
		}.Build().Output(&outputPath)

		data = append(
			data,
			pdfData,
		)
	}

	return &data, &mountPath, nil
}

func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	rand.Seed(time.Now().UnixNano())

	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}

func mkdirP(dirPath string) error {
	_, err := os.Stat(dirPath)
	if os.IsNotExist(err) {
		err = os.MkdirAll(dirPath, 0755)
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	return nil
}

func copyPDF(srcPath, destPath string) error {
	srcFile, err := ioutil.ReadFile(srcPath)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(destPath, srcFile, 0644)
	if err != nil {
		return err
	}

	return nil
}

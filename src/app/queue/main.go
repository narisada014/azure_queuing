package queue

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"os"
	"time"

	"github.com/Azure/azure-storage-queue-go/azqueue"
)

// Please set the ACCOUNT_NAME and ACCOUNT_KEY environment variables to your storage account's
// name and account key, before running the examples.
func accountInfo() (string, string) {
	return os.Getenv("ACCOUNT_NAME"), os.Getenv("ACCOUNT_KEY")
}

func main() {
	// From the Azure portal, get your Storage account's name and account key.
	accountName, accountKey := accountInfo()

	// Use your Storage account's name and key to create a credential object; this is used to access your account.
	credential, err := azqueue.NewSharedKeyCredential(accountName, accountKey)
	if err != nil {
		log.Fatal(err)
	}

	// Create a request pipeline that is used to process HTTP(S) requests and responses. It requires
	// your account credentials. In more advanced scenarios, you can configure telemetry, retry policies,
	// logging, and other options. Also, you can configure multiple request pipelines for different scenarios.
	p := azqueue.NewPipeline(credential, azqueue.PipelineOptions{})

	// From the Azure portal, get your Storage account queue service URL endpoint.
	// The URL typically looks like this:
	u, _ := url.Parse(fmt.Sprintf("https://%s.queue.core.windows.net", accountName))

	// Create an ServiceURL object that wraps the service URL and a request pipeline.
	serviceURL := azqueue.NewServiceURL(*u, p)

	// Now, you can use the serviceURL to perform various queue operations.

	// All HTTP operations allow you to specify a Go context.Context object to control cancellation/timeout.
	ctx := context.TODO() // This example uses a never-expiring context.

	// Create a URL that references a queue in your Azure Storage account.
	// This returns a QueueURL object that wraps the queue's URL and a request pipeline (inherited from serviceURL)
	queueURL := serviceURL.NewQueueURL("examplequeue") // Queue names require lowercase

	// The code below shows how to create the queue. It is common to create a queue and never delete it:
	_, err = queueURL.Create(ctx, azqueue.Metadata{})
	if err != nil {
		log.Fatal(err)
	}

	// The code below shows how a client application enqueues 2 messages into the queue:
	// Create a URL allowing you to manipulate a queue's messages.
	// This returns a MessagesURL object that wraps the queue's messages URL and a request pipeline (inherited from queueURL)
	messagesURL := queueURL.NewMessagesURL()

	// Enqueue 2 messages
	_, err = messagesURL.Enqueue(ctx, "This is message 1", time.Second*0, time.Minute)
	if err != nil {
		log.Fatal(err)
	}
	_, err = messagesURL.Enqueue(ctx, "This is message 2", time.Second*0, time.Minute)
	if err != nil {
		log.Fatal(err)
	}

	// The code below shows how a client or server can determine the approximate count of messages in the queue:
	props, err := queueURL.GetProperties(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Approximate number of messages in the queue=%d\n", props.ApproximateMessagesCount())
}
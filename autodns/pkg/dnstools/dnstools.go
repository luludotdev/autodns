package dnstools

import (
	"context"

	"github.com/bobesa/go-domain-util/domainutil"
	"github.com/libdns/libdns"
	"github.com/lolPants/autodns/autodns/pkg/iptools"
	"github.com/lolPants/autodns/autodns/pkg/logger"
)

// SetRecords sets the DNS records from an IPTools result
func SetRecords(setter libdns.RecordSetter, ip *iptools.Result, record string) error {
	if ip.Error != nil {
		logger.Stderr.Println(2, "iptools.Lookup() result has undhandled error!")
		return ip.Error
	}

	logger.Stdout.Println(2, "creating libdns records")
	records := make([]libdns.Record, 0)

	if ip.V4 != nil {
		r := libdns.Record{
			Type:  "A",
			Name:  record,
			Value: *ip.V4,
		}

		records = append(records, r)
	}

	if ip.V6 != nil {
		r := libdns.Record{
			Type:  "AAAA",
			Name:  record,
			Value: *ip.V6,
		}

		records = append(records, r)
	}

	ctx := context.Background()
	domain := domainutil.Domain(record)

	_, err := setter.SetRecords(ctx, domain, records)
	if err != nil {
		logger.Stderr.Printf(1, "dns record set failed! error: `%s`\n", err.Error())
	}

	return err
}

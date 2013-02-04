package brigade

import (
	"container/list"
	"launchpad.net/goamz/aws"
	"launchpad.net/goamz/s3"
	"log"
)

var ScanDirs *list.List
var DelDirs *list.List

var FileQueue chan string

type S3Connection struct {
	Source       *s3.S3
	Dest         *s3.S3
	SourceBucket *s3.Bucket
	DestBucket   *s3.Bucket
}

func S3Connect(t *Target) *s3.S3 {
	auth := aws.Auth{t.AccessKey, t.SecretAccessKey}
	return s3.New(auth, aws.Region{S3Endpoint: t.Server})
}

func S3Init() *S3Connection {
	s := &S3Connection{S3Connect(Config.Source), S3Connect(Config.Dest), nil, nil}

	if s.Source == nil {
		log.Fatalf("Could not connect to S3 endpoint %s", Config.Source.Server)
	}

	if s.Dest == nil {
		log.Fatalf("Could not connect to S3 endpoint %s", Config.Dest.Server)
	}

	s.SourceBucket = s.Source.Bucket(Config.Source.BucketName)
	s.DestBucket = s.Source.Bucket(Config.Dest.BucketName)

	return s
}

func (s *S3Connection) fileWorker() {
	// pull files off channel, copy with permissions
}

func InitLists() {
	ScanDirs = list.New()
  DelDirs = list.New()
}

func (s *S3Connection) CopyBucket() {

	// spawn workers
	for i := 0; i < Config.Workers; i++ {
		go fileCopier()
	}

}

func inList(input string, list []string) bool {
  for i := 0; i < len(list); i++ {
    if (input == list[i]) {
      return true
    }
  }
  return false
}

func (s *S3Connection) CopyDirectory(dir string) error {

	sourceList, err := s.SourceBucket.List(dir, "/", "", 1000)
  if err != nil {
    return err
  }

	destList, err := s.DestBucket.List(dir, "/", "", 1000)
  if err != nil {
    return err
  }

	// push subdirectories onto directory queue
  for i := 0; i < len(sourceList.CommonPrefixes); i++ {
    ScanDirs.PushBack(sourceList.CommonPrefixes[i])
  }

  // push subdirectories that no longer exist onto delete queue
  for i := 0; i < len(destList.CommonPrefixes); i++ {
    if !inList(destList.CommonPrefixes[i], sourceList.CommonPrefixes) {
      DelDirs.PushBack(destList.CommonPrefixes[i])
    }
  }

	// push changed files onto file queue

  return nil
}

func fileCopier() {
}

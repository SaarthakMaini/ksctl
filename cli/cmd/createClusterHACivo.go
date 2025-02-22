package cmd

/*
Kubesimplify
@maintainer: 	Dipankar Das <dipankardas0115@gmail.com>
				Anurag Kumar <contact.anurag7@gmail.com>
				Avinesh Tripathi <avineshtripathi1@gmail.com>
*/
import (
	log "github.com/kubesimplify/ksctl/api/logger"

	"github.com/kubesimplify/ksctl/api/civo"
	"github.com/kubesimplify/ksctl/api/utils"
	"github.com/spf13/cobra"
)

var createClusterHACivo = &cobra.Command{
	Use:   "ha-civo",
	Short: "Use to create a HA CIVO k3s cluster",
	Long: `It is used to create cluster with the given name from user. For example:

ksctl create-cluster ha-civo <arguments to civo cloud provider>
`,
	Run: func(cmd *cobra.Command, args []string) {
		isSet := cmd.Flags().Lookup("verbose").Changed
		logger := log.Logger{Verbose: true}
		if !isSet {
			logger.Verbose = false
		}
		payload := civo.CivoProvider{
			ClusterName: chcclustername,
			Region:      chcregion,
			HACluster:   true,
			Spec: utils.Machine{
				Disk:                chcnodesize,
				HAControlPlaneNodes: chcnocp,
				HAWorkerNodes:       chcnowp,
			},
		}
		err := payload.CreateCluster(logger)
		if err != nil {
			logger.Err(err.Error())
			return
		}
		logger.Info("CREATED CLUSTER")
	},
}

var (
	// c hc -> create to ha-civo
	chcregion      string
	chcclustername string
	chcnodesize    string
	chcnocp        int
	chcnowp        int
)

func init() {
	createClusterCmd.AddCommand(createClusterHACivo)
	createClusterHACivo.Flags().StringVarP(&chcnodesize, "nodeSize", "s", "g3.small", "Node size")
	createClusterHACivo.Flags().StringVarP(&chcclustername, "name", "n", "", "Cluster name")
	createClusterHACivo.Flags().StringVarP(&chcregion, "region", "r", "LON1", "Region")
	createClusterHACivo.Flags().IntVarP(&chcnocp, "control-nodes", "c", 3, "no of control plane nodes")
	createClusterHACivo.Flags().IntVarP(&chcnowp, "worker-nodes", "w", 1, "no of worker nodes")
	createClusterHACivo.Flags().BoolP("verbose", "v", true, "for verbose output")
	createClusterHACivo.MarkFlagRequired("name")
}

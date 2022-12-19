package cmd

func flags() {

	listen.Flags().StringVarP(&port, "port", "p", "0.0.0.0", "Port set to be listened")

}

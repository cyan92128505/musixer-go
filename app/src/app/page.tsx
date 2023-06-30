"use client";

import { Button, Modal, ModalBody, ModalCloseButton, ModalContent, ModalFooter, ModalHeader, ModalOverlay, useDisclosure } from "@chakra-ui/react";
import LoremIpsum from "react-lorem-ipsum";

export default function Home() {
  const { isOpen, onOpen, onClose } = useDisclosure();

  return (
    <main className="flex min-h-screen flex-col items-center justify-between p-24">
      <div>
        <Button colorScheme="blue" onClick={onOpen}>
          Hello world!
        </Button>
        <Modal isOpen={isOpen} onClose={onClose}>
          <ModalOverlay />
          <ModalContent color={"blackAlpha.900"}>
            <ModalHeader>Modal Title</ModalHeader>
            <ModalCloseButton />
            <ModalBody>
              <LoremIpsum p={2} />
            </ModalBody>

            <ModalFooter>
              <Button
                color={"blackAlpha.900"}
                mr={3}
                onClick={onClose}
              >
                Close
              </Button>
              <Button variant="ghost">Secondary Action</Button>
            </ModalFooter>
          </ModalContent>
        </Modal>
      </div>
    </main>
  );
}
